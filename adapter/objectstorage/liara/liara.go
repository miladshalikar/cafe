package liaraobjectstorage

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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
	client *s3.Client
}

func New(credential Config) *Disk {

	cfg, lErr := config.LoadDefaultConfig(context.Background(), config.WithRegion("us-west-2"))
	if lErr != nil {
		log.Println("error in load config in liara object storage :", lErr)
	}
	cfg.Credentials = aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		return aws.Credentials{
			AccessKeyID:     credential.AwsAccessKey,
			SecretAccessKey: credential.AwsSecretAccessKey,
		}, nil
	})

	cfg.BaseEndpoint = aws.String(credential.Endpoint)

	client := s3.NewFromConfig(cfg)

	return &Disk{
		config: credential,
		client: client,
	}
}
