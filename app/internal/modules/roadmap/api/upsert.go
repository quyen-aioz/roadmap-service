package roadmapapi

import (
	"context"
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"
	"roadmap/pkg/httpx"
	"time"
)

type (
	SyncRoadmapRequest struct {
		Body struct {
			Upsert []struct {
				ID        string               `json:"id"`
				Title     string               `json:"title" validate:"required"`
				Content   string               `json:"content"`
				Status    roadmapmodel.Status  `json:"status" validate:"required"`
				GroupID   roadmapmodel.GroupID `json:"group_id" validate:"required"`
				StartDate time.Time            `json:"start_date" validate:"required"`
				EndDate   time.Time            `json:"end_date" validate:"required"`
			} `json:"roadmap"`
			Delete []string `json:"delete_ids"`
		}
	}
	SyncRoadmapResponse struct {
		httpx.Paging[roadmapmodel.RoadmapDTO] `json:",inline"`
	}
)

func (h *Handler) SyncRoadmap(ctx context.Context, req *SyncRoadmapRequest) (*SyncRoadmapResponse, error) {
	roadmaps := make([]roadmapmodel.Roadmap, len(req.Body.Upsert))
	for i, r := range req.Body.Upsert {
		if !r.Status.IsValid() {
			return nil, roadmapmodel.ErrInvalidStatus
		}
		if !r.GroupID.IsValid() {
			return nil, roadmapmodel.ErrInvalidGroupID
		}

		roadmaps[i] = roadmapmodel.Roadmap{
			ID:        r.ID,
			Title:     r.Title,
			Content:   r.Content,
			Status:    r.Status,
			GroupID:   r.GroupID,
			StartDate: r.StartDate,
			EndDate:   r.EndDate,
		}
	}

	result, err := h.svc.SyncRoadmap(ctx, roadmaps, req.Body.Delete)
	if err != nil {
		return nil, err
	}

	dtos := make([]roadmapmodel.RoadmapDTO, len(result))
	for i, r := range result {
		dtos[i] = roadmapmodel.RoadmapDTO{
			ID:        r.ID,
			Title:     r.Title,
			Content:   r.Content,
			Status:    r.Status,
			GroupID:   r.GroupID,
			StartDate: r.StartDate,
			EndDate:   r.EndDate,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		}
	}

	return &SyncRoadmapResponse{
		Paging: httpx.Paging[roadmapmodel.RoadmapDTO]{
			Total: int64(len(dtos)),
			Data:  dtos,
		},
	}, nil
}
