package sdk

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
)

func GetSessionWithEnv() (*session.Session, error) {
	return session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
			AccessKeyID:     os.Getenv("AWS_ACCESS_KEY"),
			SecretAccessKey: os.Getenv("AWS_SECRET_KEY"),
		}),
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
}

func GetAWSStringValues(values []string) []*string {
	result := make([]*string, len(values))
	for _, v := range values {
		result = append(result, aws.String(v))
	}
	return result
}
