package roadmapgrouprepo

import (
	"context"
	roadmapgroupmodel "roadmap/app/internal/modules/roadmapgroup/model"

	"gorm.io/gorm"
)

type Repository interface {
	GetRoadmapGroup(ctx context.Context) ([]roadmapgroupmodel.RoadmapGroup, error)
	BulkUpsert(ctx context.Context, groups []roadmapgroupmodel.RoadmapGroup, updateColumns ...string) error
}

type SqliteRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *SqliteRepo {
	return &SqliteRepo{
		db: db,
	}
}
