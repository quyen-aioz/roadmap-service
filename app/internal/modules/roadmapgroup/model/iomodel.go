package roadmapgroupmodel

import (
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"
	"time"
)

type RoadmapGroupDTO struct {
	ID        roadmapmodel.GroupID `json:"id"`
	Name      string               `json:"name"`
	SvgURL    string               `json:"svg_url"`
	SortOrder int                  `json:"sort_order"`
	IsActive  bool                 `json:"is_active"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}
