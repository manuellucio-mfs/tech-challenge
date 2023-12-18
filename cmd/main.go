package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	awsactions "tech-challenge/aws-actions"
	"tech-challenge/helpers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(event events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {

	//Obtener valores iniciales
	fmt.Println("Obrtener valores inciales")
	paramas := awsactions.GetParam()
	var transactionByMount map[string]helpers.SimpleTransaction
	contByMount, transactionByMount := helpers.InicializarData()

	//Obtener archivo de S3
	fmt.Println("Obtener archivo de s3")
	pendingFileBuket := paramas.S3Docs.PendingTransaction
	pathFile := paramas.PathFile
	fileTransaction, err := awsactions.GetFileTransaction(pendingFileBuket, pathFile)
	if err != nil {
		fmt.Println("error to get trnasaction file")
		response = events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    map[string]string{"Content-Type": "text/plain"},
			Body:       "error to get trnasaction file",
		}
		return
	}

	//Leer el archivo y obtener la data
	/*datafile, err := filemanager.ReadFile(pathFile)
	if err != nil {
		fmt.Println("error to read file")
		return
	}*/
	fmt.Println("Obtener data")
	datafile, _ := ioutil.ReadAll(fileTransaction.Body)
	data := strings.Split(string(datafile), "\n")

	//Almacenar la data
	var contDebit, contCredit int
	var sumDebit, sumCredit float32
	var creditTransactions, debitTransactions []helpers.Transaction
	var typeTransaction bool

	fmt.Println("Iniciar procesamiento de la data")
	for position, row := range data {
		if position > 0 {
			dataRow := strings.Split(row, ",")

			valueTransaction := strings.TrimSpace(dataRow[2][1:])
			value, e := strconv.ParseFloat(valueTransaction, 32)
			if err != nil {
				fmt.Println("errot to convert value", e)
			}

			transaction := helpers.Transaction{
				IdTransaction: dataRow[0],
				Date:          dataRow[1],
				Value:         dataRow[2][1:],
			}

			//Realizar sumarizacion de tipo de transaccion
			if strings.Contains(dataRow[2], "-") {

				sumCredit = sumCredit + float32(value)
				typeTransaction = false
				contCredit++

				creditTransactions = append(creditTransactions, transaction)
			} else {

				sumDebit = sumDebit + float32(value)
				typeTransaction = true
				contDebit++

				debitTransactions = append(debitTransactions, transaction)
			}

			//Conteo de transacciones por mes
			valueDate := strings.Split(dataRow[1], "/")
			mount, e := strconv.Atoi(valueDate[0])
			if err != nil {
				fmt.Println("error to conver mount to int:", e)
				response = events.APIGatewayProxyResponse{
					StatusCode: 500,
					Headers:    map[string]string{"Content-Type": "text/plain"},
					Body:       "error to conver mount to int:",
				}
				return
			}
			//Sumarizacion de transacion por mes
			transactionByMount = helpers.TransactionsByMount(mount, float32(value), typeTransaction, transactionByMount)
			contByMount = helpers.CountTransactionsByMount(mount, contByMount)
		}
	}

	//Calcular total y promedio de cuenta
	fmt.Println("Realizar calculos")
	totalBalance := sumDebit - sumCredit
	averageDebit := sumDebit / float32(contDebit)
	averageCredit := sumCredit / float32(contCredit)

	//Guardar transaccion de BD
	fmt.Println("Guardar transaccion de BD")
	filestruct := strings.Split(pathFile, ".")
	account := helpers.Account{
		IdAccount:          filestruct[0],
		TotalBalance:       totalBalance,
		AverageCredit:      averageCredit,
		AverageDebit:       averageDebit,
		CreditTransactions: creditTransactions,
		DebitTransactions:  debitTransactions,
	}
	awsactions.SaveTransactions(paramas.TableName, account)

	//Guardar transaccion en S3
	fmt.Println("Guardar transaccion en s3")
	err = awsactions.UploadFile(datafile, paramas.S3Docs.ProcessedTransaction, pathFile)
	if err != nil {
		fmt.Println("error to upload file on s3", err)
		response = events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    map[string]string{"Content-Type": "text/plain"},
			Body:       "error to upload file on s3",
		}
		return
	}

	//Imprinir resultados
	fmt.Println("Imprimir resultados")
	msgBalance := fmt.Sprintf("Good day dear user with billing number %s \n We sent you the credit and debit balance information \n\n", filestruct[0])
	msgBalance = msgBalance + fmt.Sprintf("Total Balance: %f \n", totalBalance)
	fmt.Println("Total Balance: ", totalBalance)
	for k, v := range contByMount {
		if v > 0 {
			fmt.Println("Number transaction in", k, ":", v)
			msgBalance = msgBalance + fmt.Sprintf("Number transaction in %s: %d \n", k, v)
		}
	}
	fmt.Println("Average Debit: ", averageDebit)
	msgBalance = msgBalance + fmt.Sprintf("Average Debit: %f \n", averageDebit)
	fmt.Println("Average Credit: ", averageCredit)
	msgBalance = msgBalance + fmt.Sprintf("Average Credi: %f \n\n", averageCredit)

	for k, v := range transactionByMount {
		if v.CantTranCredit > 0 || v.CantTranDebit > 0 {
			fmt.Println("Number transaction in", k, ":", v.CantTranCredit+v.CantTranDebit)
			msgBalance = msgBalance + fmt.Sprintf("Number transaction in %s: %d \n", k, v.CantTranCredit+v.CantTranDebit)
			totalBalanceByMount := v.SumarDebit - v.SumaryCredit
			msgBalance = msgBalance + fmt.Sprintf("total balance: %f \n", totalBalanceByMount)
			if v.CantTranCredit > 0 {
				averageCreditByMount := v.SumaryCredit / float32(v.CantTranCredit)
				msgBalance = msgBalance + fmt.Sprintf("Average credit: %f \n", averageCreditByMount)
			}
			if v.CantTranDebit > 0 {
				averageDebitByMount := v.SumarDebit / float32(v.CantTranDebit)
				msgBalance = msgBalance + fmt.Sprintf("Average debit: %f \n", averageDebitByMount)
			}
			msgBalance = msgBalance + "\n"
		}
	}
	msgBalance = msgBalance + "We send cordial greetings wishing happy holidays"

	//Enviar Email
	err = awsactions.SendEmail(msgBalance, paramas.SnsArn)
	if err != nil {
		fmt.Println("error to send email", err)
	}

	//If all is OK up to this point send a successful response
	response = events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "text/plain"},
		Body:       "Successful Request",
	}
	return
}
