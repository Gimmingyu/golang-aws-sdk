package sdk

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"os"
)

func NewS3Client(sess *session.Session) *s3.S3 {
	return s3.New(sess)
}

func PutObject(svc *s3.S3, bucket, key, path string) (*s3.PutObjectOutput, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Panicf("panic occurred while opening file %s: %v", path, err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	return svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})
}

func ListObjects(svc *s3.S3, bucket string) (*s3.ListObjectsOutput, error) {
	return svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucket)})
}

func GetObject(svc *s3.S3, bucket, key string) (*s3.GetObjectOutput, error) {
	return svc.GetObject(&s3.GetObjectInput{Bucket: aws.String(bucket), Key: aws.String(key)})
}
