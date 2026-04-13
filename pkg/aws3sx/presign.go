package aws3sx

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	defaultPresignExpiry = 5 * time.Minute
)

type (
	PresignUploadRequest struct {
		Key         string
		ContentType string
		Expiry      time.Duration
	}
	PresignUploadResponse struct {
		UploadURL string
		PublicURL string
	}
)

func (c *Client) PresignUpload(ctx context.Context, req PresignUploadRequest) (PresignUploadResponse, error) {
	if req.Key == "" {
		return PresignUploadResponse{}, fmt.Errorf("key is required")
	}

	if req.ContentType == "" {
		return PresignUploadResponse{}, fmt.Errorf("content type is required")
	}

	expiry := defaultPresignExpiry
	if req.Expiry > 0 {
		expiry = req.Expiry
	}

	presigned, err := c.presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(c.bucket),
		Key:         aws.String(req.Key),
		ContentType: aws.String(req.ContentType),
	}, s3.WithPresignExpires(expiry))
	if err != nil {
		return PresignUploadResponse{}, fmt.Errorf("failed to presign upload: %w", err)
	}

	publicURL := c.buildPublicURL(req.Key)

	return PresignUploadResponse{
		UploadURL: presigned.URL,
		PublicURL: publicURL,
	}, nil
}

func (c *Client) buildPublicURL(key string) string {
	return fmt.Sprintf("%s/s/%s/%s/%s", c.cfg.PublicEndpoint, c.cfg.AccessKey, c.cfg.Bucket, key)
}
