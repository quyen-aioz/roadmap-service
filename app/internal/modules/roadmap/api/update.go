package roadmapapi

import (
	"context"
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"
	"roadmap/pkg/humax"
	"time"
)

type (
	CreateRoadmapRequest struct {
		Body struct {
			Title     string              `json:"title" validate:"required"`
			Content   string              `json:"content" validate:"required"`
			Status    roadmapmodel.Status `json:"status" validate:"required"`
			GroupID   string              `json:"group_id" validate:"required"`
			StartDate time.Time           `json:"start_date" validate:"required"`
			EndDate   time.Time           `json:"end_date" validate:"required"`
		}
	}
	CreateRoadmapResponse struct {
		ID string `json:"id"`
	}
)

func (h *Handler) CreateRoadmap(ctx context.Context, req *CreateRoadmapRequest) (*CreateRoadmapResponse, error) {
	status := req.Body.Status

	if !status.IsValid() {
		return nil, roadmapmodel.ErrInvalidStatus
	}

	roadmapID, err := h.svc.CreateRoadmap(ctx, &roadmapmodel.Roadmap{
		Title:     req.Body.Title,
		Content:   req.Body.Content,
		Status:    status,
		GroupID:   req.Body.GroupID,
		StartDate: req.Body.StartDate,
		EndDate:   req.Body.EndDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &CreateRoadmapResponse{
		ID: roadmapID,
	}, nil
}

type (
	UpdateRoadmapRequest struct {
		ID   string `path:"id" validate:"required"`
		Body struct {
			Title     *string              `json:"title,omitempty"`
			Content   *string              `json:"content,omitempty"`
			Status    *roadmapmodel.Status `json:"status,omitempty"`
			GroupID   *string              `json:"group_id,omitempty"`
			StartDate *time.Time           `json:"start_date,omitempty"`
			EndDate   *time.Time           `json:"end_date,omitempty"`
		}
	}
)

func (h *Handler) UpdateRoadmap(ctx context.Context, req *UpdateRoadmapRequest) (*roadmapmodel.RoadmapDTO, error) {
	status := req.Body.Status

	if status != nil && !status.IsValid() {
		return nil, roadmapmodel.ErrInvalidStatus
	}

	roadmap, err := h.svc.UpdateRoadmap(ctx, req.ID, roadmapmodel.UpdateRoadmapReq{
		Title:     req.Body.Title,
		Content:   req.Body.Content,
		Status:    status,
		GroupID:   req.Body.GroupID,
		StartDate: req.Body.StartDate,
		EndDate:   req.Body.EndDate,
	})

	if err != nil {
		return nil, err
	}

	return &roadmapmodel.RoadmapDTO{
		ID:        roadmap.ID,
		Title:     roadmap.Title,
		Content:   roadmap.Content,
		Status:    roadmap.Status,
		GroupID:   roadmap.GroupID,
		StartDate: roadmap.StartDate,
		EndDate:   roadmap.EndDate,
		CreatedAt: roadmap.CreatedAt,
		UpdatedAt: roadmap.UpdatedAt,
	}, nil
}

type DeleteRoadmapRequest struct {
	ID string `path:"id" validate:"required"`
}

func (h *Handler) DeleteRoadmap(ctx context.Context, req *DeleteRoadmapRequest) (*humax.Empty, error) {
	err := h.svc.DeleteRoadmap(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &humax.Empty{}, nil
}
