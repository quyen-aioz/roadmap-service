package roadmaprepo

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"
)

type Repository interface {
	GetRoadmap(ctx context.Context) ([]roadmapmodel.Roadmap, error)
	Create(ctx context.Context, roadmap *roadmapmodel.Roadmap) (string, error)
	BulkUpsert(ctx context.Context, roadmaps []roadmapmodel.Roadmap) error
	BulkDelete(ctx context.Context, deleteIDs []string) error
	Sync(ctx context.Context, roadmaps []roadmapmodel.Roadmap, deleteIDs []string) ([]roadmapmodel.Roadmap, error)
	FindOne(ctx context.Context, q roadmapmodel.FindQueryBuilder) (roadmapmodel.Roadmap, error)
	Update(ctx context.Context, id string, u roadmapmodel.RoadmapUpdateBuilder) (string, error)
	Delete(ctx context.Context, id string) error
}

type SqliteRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *SqliteRepo {
	return &SqliteRepo{
		db: db,
	}
}

func GenerateHexID() string {
	b := make([]byte, 8) // 8 bytes = 16 hex characters
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}
