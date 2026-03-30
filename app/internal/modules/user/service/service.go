package userservice

import (
	"fmt"
	userrepo "roadmap/app/internal/modules/user/repo"
	"roadmap/pkg/sqlitex"
)

type Service struct {
	repo userrepo.Repository
}

func NewWithRepo(repo userrepo.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func New() (*Service, error) {
	db, err := sqlitex.Get()
	if err != nil {
		return nil, fmt.Errorf("sqlite: %w", err)
	}

	return &Service{
		repo: userrepo.New(db),
	}, nil
}
