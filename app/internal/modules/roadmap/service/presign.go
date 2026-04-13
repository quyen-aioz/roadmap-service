package roadmapservice

import (
	"context"
	"fmt"
	"path/filepath"
	"roadmap/app/internal/core/serverconfig"
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"
	"roadmap/pkg/aws3sx"
	"time"
)

func (svc *Service) PresignUpload(ctx context.Context, req roadmapmodel.PresignUploadReq) (roadmapmodel.PresignUploadResp, error) {
	key := buildRoadmapThumbnailKey(req.TaskID, req.ContentType)

	s3Client, err := aws3sx.Get()
	if err != nil {
		return roadmapmodel.PresignUploadResp{}, fmt.Errorf("presign upload: get s3 client: %w", err)
	}

	result, err := s3Client.PresignUpload(ctx, aws3sx.PresignUploadRequest{
		Key:         key,
		ContentType: req.ContentType,
	})
	if err != nil {
		return roadmapmodel.PresignUploadResp{}, fmt.Errorf("presign upload: %w", err)
	}

	return roadmapmodel.PresignUploadResp{
		UploadURL: result.UploadURL,
		PublicURL: result.PublicURL,
	}, nil
}

// Format: {pathFolder}/{taskID}_{unixTimestamp}{ext}
func buildRoadmapThumbnailKey(taskID, contentType string) string {
	conf := serverconfig.Get()
	ext := aws3sx.ContentTypeToExt(contentType)

	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("%s_%d%s", taskID, timestamp, ext)
	return filepath.Join(conf.W3Storage.PathFolder, filename)
}
