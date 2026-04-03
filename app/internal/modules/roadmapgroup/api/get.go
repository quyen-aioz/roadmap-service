package roadmapgroupapi

import (
	"context"
	roadmapgroupmodel "roadmap/app/internal/modules/roadmapgroup/model"
	"roadmap/pkg/httpx"
	"roadmap/pkg/humax"
)

type GetRoadmapGroupResp struct {
	httpx.Paging[roadmapgroupmodel.RoadmapGroupDTO] `json:",inline"`
}

func (h *Handler) GetRoadmapGroup(ctx context.Context, _ *humax.Empty) (*GetRoadmapGroupResp, error) {
	groups, err := h.svc.GetRoadmapGroup(ctx)
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

	return &GetRoadmapGroupResp{
		Paging: httpx.Paging[roadmapgroupmodel.RoadmapGroupDTO]{
			Total: int64(len(groups)),
			Data:  dtos,
		},
	}, nil
}
