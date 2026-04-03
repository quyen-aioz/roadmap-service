package roadmaprepo

import (
	"context"
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"

	"gorm.io/gorm"
)

type Repository interface {
	GetRoadmap(ctx context.Context) ([]roadmapmodel.Roadmap, error)
	GetRoadmapContent(ctx context.Context) (roadmapmodel.RoadmapContent, error)
	UpdateRoadmapContent(ctx context.Context, roadmapContent roadmapmodel.UpdateRoadmapContentReq) error
	Create(ctx context.Context, roadmap *roadmapmodel.Roadmap) (string, error)
	BulkUpsert(ctx context.Context, roadmaps []roadmapmodel.Roadmap) error
	BulkDelete(ctx context.Context, deleteIDs []string) error
	Sync(ctx context.Context, roadmaps []roadmapmodel.Roadmap, deleteIDs []string) ([]roadmapmodel.Roadmap, error)
	FindOne(ctx context.Context, q roadmapmodel.FindQueryBuilder) (roadmapmodel.Roadmap, error)
	Update(ctx context.Context, id string, u roadmapmodel.RoadmapUpdateBuilder) (string, error)
	Delete(ctx context.Context, id string) error
}

type SqliteRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *SqliteRepo {
	return &SqliteRepo{
		db: db,
	}
}
