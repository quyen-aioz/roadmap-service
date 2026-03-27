package roadmapapi

import (
	"context"
)

type (
	CreateRoadmapRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	CreateRoadmapResponse struct {
		ID string `json:"id"`
	}
)

func (h *Handler) CreateRoadmap(_ context.Context, _ *CreateRoadmapRequest) (*CreateRoadmapResponse, error) {
	return nil, nil
}
