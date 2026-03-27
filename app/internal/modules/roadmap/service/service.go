package roadmapservice

import (
	"fmt"
	roadmaprepo "roadmap/app/internal/modules/roadmap/repo"
	"roadmap/pkg/sqlitex"
)

type Service struct {
	repo roadmaprepo.Repository
}

func New() (*Service, error) {
	db, err := sqlitex.Get()
	if err != nil {
		return nil, fmt.Errorf("sqlite: %w", err)
	}

	return &Service{
		repo: roadmaprepo.New(db),
	}, nil
}
