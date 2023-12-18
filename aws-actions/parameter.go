package awsactions

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	_ "log"

	"github.com/aws/aws-sdk-go/service/ssm"
	_ "github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

type ParamStruct struct {
	S3Docs struct {
		PendingTransaction   string `json:"pendingTransaction"`
		ProcessedTransaction string `json:"processedTransaction"`
	} `json:"s3docs"`
	PathFile  string `json:"pathFile"`
	TableName string `json:"tableName"`
	SnsArn    string `json:"snsArn"`
}

func GetParam() ParamStruct {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String("us-east-1")},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		panic(err)
	}

	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion("us-east-1"))
	param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("tech-challenge"),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	value := *param.Parameter.Value
	var parameters ParamStruct
	rdr := strings.NewReader(value)
	if err := json.NewDecoder(rdr).Decode(&parameters); err != nil {
		fmt.Printf("error deserializing JSON: %v", err)

	}

	return parameters
}
