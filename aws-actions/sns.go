package awsactions

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func SendEmail(msg, topic string) error {
	msgPtr := flag.String("m", msg, "The message to send to the subscribed users of the topic")
	topicPtr := flag.String("t", topic, "The ARN of the topic to which the user subscribes")

	flag.Parse()

	if *msgPtr == "" || *topicPtr == "" {
		return fmt.Errorf("error to recive information, verify your data")
	}

	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	if err != nil {
		return err
	}
	svc := sns.New(sess)

	result, err := svc.Publish(&sns.PublishInput{
		Message:  msgPtr,
		TopicArn: topicPtr,
	})
	if err != nil {
		fmt.Println("error to publish message", err)
		return err
	}

	fmt.Println(*result.MessageId)
	return nil
}
