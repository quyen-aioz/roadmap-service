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
			ID:            r.ID,
			Title:         r.Title,
			Content:       r.Content,
			Status:        r.Status,
			GroupID:       r.GroupID,
			CTALabel:      r.CTALabel,
			CTALink:       r.CTALink,
			StartDate:     r.StartDate,
			EndDate:       r.EndDate,
			ThumbnailURL:  r.ThumbnailURL,
			ThumbnailType: r.ThumbnailType,
			CreatedAt:     r.CreatedAt,
			UpdatedAt:     r.UpdatedAt,
		})
	}

	return &GetRoadmapResp{
		Paging: httpx.Paging[roadmapmodel.RoadmapDTO]{
			Data:  roadmapResp,
			Total: int64(len(roadmapResp)),
		},
	}, nil
}

type GetRoadmapContentResp struct {
	roadmapmodel.RoadmapContentDTO `json:",inline"`
}

func (h *Handler) GetRoadmapContent(ctx context.Context, _ *humax.Empty) (*GetRoadmapContentResp, error) {
	roadmapContent, err := h.svc.GetRoadmapContent(ctx)
	if err != nil {
		return nil, err
	}

	return &GetRoadmapContentResp{
		RoadmapContentDTO: roadmapmodel.RoadmapContentDTO{
			ID:          roadmapContent.ID,
			Title:       roadmapContent.Title,
			Description: roadmapContent.Description,
			Content:     roadmapContent.Content,
			CreatedAt:   roadmapContent.CreatedAt,
			UpdatedAt:   roadmapContent.UpdatedAt,
		},
	}, nil
}
