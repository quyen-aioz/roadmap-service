package roadmapgroupmodel

import (
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"
	"time"

	"gorm.io/gorm"
)

type RoadmapGroup struct {
	ID        roadmapmodel.GroupID `gorm:"primaryKey"`
	Name      string
	SvgURL    string
	SortOrder int  `gorm:"index"`
	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (RoadmapGroup) TableName() string {
	return "roadmap_group"
}
