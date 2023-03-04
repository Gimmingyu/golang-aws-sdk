package amazon

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func NewSTSClient(sess *session.Session) *sts.STS {
	return sts.New(sess)
}

func AssumeRole(svc *sts.STS, roleArn, sessionName, policy string, duration int64) (*sts.AssumeRoleOutput, error) {
	return svc.AssumeRole(&sts.AssumeRoleInput{
		DurationSeconds: aws.Int64(duration),
		RoleArn:         aws.String(roleArn),
		RoleSessionName: aws.String(sessionName),
	})
}

func GetCredentialsFromResult(output *sts.AssumeRoleOutput) (string, string, string) {
	return aws.StringValue(output.Credentials.AccessKeyId), aws.StringValue(output.Credentials.SecretAccessKey), aws.StringValue(output.Credentials.SessionToken)
}
