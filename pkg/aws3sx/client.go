package aws3sx

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var instance *Client

type Client struct {
	s3            *s3.Client
	bucket        string
	cfg           Config
	presignClient *s3.PresignClient
}

type Config struct {
	AccessKey      string
	SecretKey      string
	Endpoint       string
	PublicEndpoint string
	Bucket         string
	Region         string
}

func Init(cfg Config) error {
	if cfg.AccessKey == "" || cfg.SecretKey == "" {
		return fmt.Errorf("aws s3: missing access key or secret key")
	}

	if cfg.Endpoint == "" {
		return fmt.Errorf("aws s3: missing endpoint")
	}

	if cfg.Bucket == "" {
		return fmt.Errorf("aws s3: missing bucket")
	}

	region := cfg.Region
	if region == "" {
		region = "us-east-1" // default for S3-compatible storage
	}

	awsCfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretKey, ""),
		),
	)
	if err != nil {
		return fmt.Errorf("aws s3: load config: %w", err)
	}

	s3Client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(cfg.Endpoint)
		o.UsePathStyle = true // required for most S3-compatible storage (path-style: endpoint/bucket/key)
	})

	instance = &Client{
		s3:            s3Client,
		bucket:        cfg.Bucket,
		cfg:           cfg,
		presignClient: s3.NewPresignClient(s3Client),
	}

	return nil

}

func Get() (*Client, error) {
	if instance == nil {
		return nil, fmt.Errorf("s3x: client not initialized, call Init first")
	}
	return instance, nil
}
