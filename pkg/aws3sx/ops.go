package aws3sx

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Object struct {
	Key          string
	LastModified time.Time
	Size         int64
}

type ListObjectsInput struct {
	Prefix string
}

// ListObjects returns all objects under the given prefix in the bucket.
// Handles pagination internally — always returns the full list.
func (c *Client) ListObjects(ctx context.Context, input ListObjectsInput) ([]S3Object, error) {
	var results []S3Object
	var continuationToken *string

	for {
		resp, err := c.s3.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
			Bucket:            aws.String(c.bucket),
			Prefix:            aws.String(input.Prefix),
			ContinuationToken: continuationToken,
		})
		if err != nil {
			return nil, fmt.Errorf("aws3sx: list objects: %w", err)
		}

		for _, obj := range resp.Contents {
			results = append(results, S3Object{
				Key:          aws.ToString(obj.Key),
				LastModified: aws.ToTime(obj.LastModified),
				Size:         aws.ToInt64(obj.Size),
			})
		}

		// no more pages
		if !aws.ToBool(resp.IsTruncated) {
			break
		}

		// next page
		continuationToken = resp.NextContinuationToken
	}

	return results, nil
}

// DeleteObject deletes a single object by key from the bucket.
func (c *Client) DeleteObject(ctx context.Context, key string) error {
	_, err := c.s3.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("aws3sx: delete object %q: %w", key, err)
	}
	return nil
}

// ExtractKeyFromURL parses the S3 object key out of a full public AIOZ URL.
// URL format: {endpoint}/s/{accessKey}/{bucket}/{key}
// Example: https://api-pub-storage.attoaioz.cyou/s/ACCESSKEY/test-roadmap/roadmap/tasks/thumbnails/abc_123.jpg
// Returns: roadmap/tasks/thumbnails/abc_123.jpg
func (c *Client) ExtractKeyFromURL(publicURL string) string {
	raw := strings.TrimSpace(publicURL)
	if raw == "" {
		return ""
	}

	// handle full URL
	if strings.Contains(raw, "://") {
		parsed, err := url.Parse(raw)
		if err != nil {
			return ""
		}
		return extractKeyFromPath(parsed.Path, c.cfg.Bucket)
	}

	// handle raw path/key (with optional query/fragment)
	if idx := strings.IndexAny(raw, "?#"); idx >= 0 {
		raw = raw[:idx]
	}
	return extractKeyFromPath(raw, c.cfg.Bucket)
}

func extractKeyFromPath(path, bucket string) string {
	cleanPath := strings.TrimSpace(path)
	if cleanPath == "" {
		return ""
	}

	cleanPath = strings.Trim(cleanPath, "/")
	if cleanPath == "" {
		return ""
	}

	decoded, err := url.PathUnescape(cleanPath)
	if err == nil {
		cleanPath = decoded
	}

	// allow input without scheme, e.g. "api-pub-storage.../s/ACCESS/bucket/key"
	parts := strings.Split(cleanPath, "/")
	if len(parts) > 1 && strings.Contains(parts[0], ".") {
		cleanPath = strings.Join(parts[1:], "/")
		parts = strings.Split(cleanPath, "/")
	}

	trimmedBucket := strings.Trim(bucket, "/")

	// format: /s/{accessKey}/{bucket}/{key}
	if len(parts) >= 4 && parts[0] == "s" && parts[2] == trimmedBucket {
		return strings.Join(parts[3:], "/")
	}

	// format: /{bucket}/{key}
	if trimmedBucket != "" {
		if cleanPath == trimmedBucket {
			return ""
		}
		prefix := trimmedBucket + "/"
		if strings.HasPrefix(cleanPath, prefix) {
			return strings.TrimPrefix(cleanPath, prefix)
		}
	}

	// already key-only path
	return cleanPath
}
