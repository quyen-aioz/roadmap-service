package roadmapapi

import (
	"context"
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"
	"roadmap/pkg/aws3sx"
)

type (
	PresignUploadRequest struct {
		Body struct {
			TaskID      string `json:"task_id" validate:"required"`
			ContentType string `json:"content_type" validate:"required"`
		}
	}
	PresignUploadResponse struct {
		UploadURL   string           `json:"upload_url"`
		PublicURL   string           `json:"public_url"`
		MediaType   aws3sx.MediaType `json:"media_type"`
		ContentType string           `json:"content_type"`
	}
)

func (h *Handler) PresignUpload(ctx context.Context, req *PresignUploadRequest) (*PresignUploadResponse, error) {
	if !aws3sx.IsAllowedContentType(req.Body.ContentType) {
		return nil, roadmapmodel.ErrUnsupportedMediaType
	}

	mediaType, err := aws3sx.DetectMediaType(req.Body.ContentType)
	if err != nil {
		return nil, roadmapmodel.ErrUnsupportedMediaType
	}

	result, err := h.svc.PresignUpload(ctx, roadmapmodel.PresignUploadReq{
		TaskID:      req.Body.TaskID,
		ContentType: req.Body.ContentType,
	})
	if err != nil {
		return nil, err
	}

	return &PresignUploadResponse{
		UploadURL:   result.UploadURL,
		PublicURL:   result.PublicURL,
		MediaType:   mediaType,
		ContentType: req.Body.ContentType,
	}, nil
}
