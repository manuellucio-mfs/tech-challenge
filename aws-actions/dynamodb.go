package awsactions

import (
	"context"
	"fmt"
	"tech-challenge/helpers"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func SaveTransactions(tableName string, transaction helpers.Account) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)
	out, err := svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"idAccount":     &types.AttributeValueMemberS{Value: transaction.IdAccount},
			"totalBalance":  &types.AttributeValueMemberS{Value: fmt.Sprintf("%f", transaction.TotalBalance)},
			"averageCredit": &types.AttributeValueMemberS{Value: fmt.Sprintf("%f", transaction.AverageCredit)},
			"averageDebit":  &types.AttributeValueMemberS{Value: fmt.Sprintf("%f", transaction.AverageDebit)},
			//"creditTransactions": &types.AttributeValueMemberL{Value: transaction.CreditTransactions},
			//"debitTransactions":  &types.AttributeValueMemberSS{Value: transaction.DebitTransactions},
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(out.Attributes)
}
