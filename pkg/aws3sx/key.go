package aws3sx

import (
	"fmt"
	"strings"
)

type MediaType string

const (
	MediaTypeImage MediaType = "image"
	MediaTypeGIF   MediaType = "gif"
	MediaTypeVideo MediaType = "video"
)

func DetectMediaType(contentType string) (MediaType, error) {
	switch {
	case contentType == "image/gif":
		return MediaTypeGIF, nil
	case strings.HasPrefix(contentType, "image/"):
		return MediaTypeImage, nil
	case strings.HasPrefix(contentType, "video/"):
		return MediaTypeVideo, nil
	default:
		return "", fmt.Errorf("aws 3s: unsupported content type: %s", contentType)
	}
}

var AllowedContentTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/webp": true,
	"image/gif":  true,
	"video/mp4":  true,
	"video/webm": true,
}

func IsAllowedContentType(contentType string) bool {
	return AllowedContentTypes[contentType]
}

func ContentTypeToExt(contentType string) string {
	m := map[string]string{
		"image/jpeg": ".jpg",
		"image/png":  ".png",
		"image/webp": ".webp",
		"image/gif":  ".gif",
		"video/mp4":  ".mp4",
		"video/webm": ".webm",
	}
	if ext, ok := m[contentType]; ok {
		return ext
	}
	return ""
}
