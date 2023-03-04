package amazon

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type (
	Attributes struct {
		Key         string
		DataType    string
		StringValue string
	}
)

func NewAttributes(key string, dataType string, stringValue string) *Attributes {
	return &Attributes{Key: key, DataType: dataType, StringValue: stringValue}
}

func NewSQSClient(sess *session.Session) *sqs.SQS {
	return sqs.New(sess)
}

func SendSimpleMessage(svc *sqs.SQS, queueUrl, messageBody string) (*sqs.SendMessageOutput, error) {
	return svc.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(messageBody),
		QueueUrl:    aws.String(queueUrl),
	})
}

func SendComplexMessage(svc *sqs.SQS, queueUrl, messageBody string, attr []Attributes) {
	svc.SendMessage(&sqs.SendMessageInput{
		QueueUrl:          aws.String(queueUrl),
		MessageAttributes: GetMessageAttributes(attr),
		MessageBody:       aws.String(messageBody),
	})
}

func GetMessageAttributes(attr []Attributes) map[string]*sqs.MessageAttributeValue {
	result := make(map[string]*sqs.MessageAttributeValue)
	for _, v := range attr {
		result[v.Key] = &sqs.MessageAttributeValue{
			DataType:    aws.String(v.DataType),
			StringValue: aws.String(v.StringValue),
		}
	}
	return result
}

func Receive(svc *sqs.SQS, queueUrl string, maxNumber int64, attributes []string) (*sqs.ReceiveMessageOutput, error) {
	return svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames:        GetAWSStringValues(attributes),
		MaxNumberOfMessages:   aws.Int64(maxNumber),
		MessageAttributeNames: GetAWSStringValues(attributes),
		QueueUrl:              aws.String(queueUrl),
	})
}
