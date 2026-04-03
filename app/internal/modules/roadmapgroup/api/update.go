package roadmapgroupapi

import (
	"context"
	roadmapgroupmodel "roadmap/app/internal/modules/roadmapgroup/model"
	"roadmap/pkg/httpx"
)

type (
	ReorderGroupRequest struct {
		Body struct {
			OrderedIDs []string `json:"ordered_ids"`
		}
	}
	ReorderGroupResponse struct {
		httpx.Paging[roadmapgroupmodel.RoadmapGroupDTO] `json:",inline"`
	}
)

func (h *Handler) ReorderGroup(ctx context.Context, req *ReorderGroupRequest) (*ReorderGroupResponse, error) {
	groups, err := h.svc.ReorderGroup(ctx, req.Body.OrderedIDs)
	if err != nil {
		return nil, err
	}

	dtos := make([]roadmapgroupmodel.RoadmapGroupDTO, len(groups))
	for i, g := range groups {
		dtos[i] = roadmapgroupmodel.RoadmapGroupDTO{
			ID:        g.ID,
			Name:      g.Name,
			SvgURL:    g.SvgURL,
			SortOrder: g.SortOrder,
			IsActive:  g.IsActive,
			CreatedAt: g.CreatedAt,
			UpdatedAt: g.UpdatedAt,
		}
	}

	return &ReorderGroupResponse{
		Paging: httpx.Paging[roadmapgroupmodel.RoadmapGroupDTO]{
			Total: int64(len(groups)),
			Data:  dtos,
		},
	}, nil
}
