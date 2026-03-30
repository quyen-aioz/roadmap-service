package userrepo

import (
	"context"
	usermodel "roadmap/app/internal/modules/user/model"

	"gorm.io/gorm"
)

type Repository interface {
	FindOne(ctx context.Context, q usermodel.FindQueryBuilder) (usermodel.User, error)
	Create(ctx context.Context, req usermodel.CreateUserReq) (usermodel.User, error)
}

type SqliteRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *SqliteRepo {
	return &SqliteRepo{
		db: db,
	}
}
