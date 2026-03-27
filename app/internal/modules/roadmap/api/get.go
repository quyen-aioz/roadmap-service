package roadmapapi

import (
	"context"
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"
	"roadmap/pkg/httpx"
	"roadmap/pkg/humax"
)

type (
	GetRoadmapResp struct {
		httpx.Paging[roadmapmodel.RoadmapDTO] `json:",inline"`
	}
)

func (h *Handler) GetRoadmap(ctx context.Context, _ *humax.Empty) (*GetRoadmapResp, error) {
	roadmap, err := h.svc.GetRoadmap(ctx)
	if err != nil {
		return nil, err
	}

	roadmapResp := make([]roadmapmodel.RoadmapDTO, 0, len(roadmap))
	for _, r := range roadmap {
		roadmapResp = append(roadmapResp, roadmapmodel.RoadmapDTO{
			ID:          r.ID,
			Name:        r.Name,
			Description: r.Description,
			CreatedAt:   r.CreatedAt,
			UpdatedAt:   r.UpdatedAt,
		})
	}

	return &GetRoadmapResp{
		Paging: httpx.Paging[roadmapmodel.RoadmapDTO]{
			Data:  roadmapResp,
			Total: int64(len(roadmapResp)),
		},
	}, nil
}
