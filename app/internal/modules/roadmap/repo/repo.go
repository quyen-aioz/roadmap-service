package roadmaprepo

import (
	"context"
	"database/sql"
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"
)

type Repository interface {
	GetRoadmap(ctx context.Context) ([]roadmapmodel.Roadmap, error)
}

type SqliteRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *SqliteRepo {
	return &SqliteRepo{
		db: db,
	}
}
