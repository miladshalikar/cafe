package liaraobjectstorage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

type Config struct {
	AwsAccessKey       string `koanf:"access_key"`
	AwsSecretAccessKey string `koanf:"secret_key"`
	Endpoint           string `koanf:"endpoint_url"`
	BucketName         string `koanf:"bucket_name"`
}

type Disk struct {
	config Config
	client *s3.S3
}

func New(credential Config) *Disk {

	sess, err := session.NewSession(&aws.Config{
		Endpoint:    aws.String(credential.Endpoint),
		Region:      aws.String("us-east-1"), // Region is required but can be arbitrary for S3-compatible storage
		Credentials: credentials.NewStaticCredentials(credential.AwsAccessKey, credential.AwsSecretAccessKey, ""),
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}

	// Create S3 client
	s3Client := s3.New(sess)

	return &Disk{
		config: credential,
		client: s3Client,
	}

}
