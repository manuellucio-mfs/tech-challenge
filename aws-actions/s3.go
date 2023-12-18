package awsactions

import (
	"bytes"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func GetFileTransaction(bucket, fileKey string) (*s3.GetObjectOutput, error) {
	//Crear session
	s3Client := s3.New(session.New(), &aws.Config{
		Region: aws.String("us-east-1"),
	})

	//Obtener objeto
	result, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileKey),
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println(result)

	return result, nil
}

func UploadFile(filename []byte, processedBucket, keyName string) error {
	//Create session
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	if err != nil {
		fmt.Println("There was an error starting a new session")
		return err
	}
	uploader := s3manager.NewUploader(sess)

	//Upload File in Bucket By Country
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(processedBucket),
		Key:    aws.String(keyName),
		Body:   bytes.NewReader(filename),
	})
	if err != nil {
		fmt.Println(fmt.Errorf("failed to upload file in country bucket, %v", err))
		return err
	}

	fmt.Println(result)
	return nil
}
