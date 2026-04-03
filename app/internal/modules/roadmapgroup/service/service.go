package roadmapgroupservice

import (
	"fmt"
	roadmapgrouprepo "roadmap/app/internal/modules/roadmapgroup/repo"
	"roadmap/pkg/sqlitex"
)

type Service struct {
	repo roadmapgrouprepo.Repository
}

func New() (*Service, error) {
	db, err := sqlitex.Get()
	if err != nil {
		return nil, fmt.Errorf("sqlite: %w", err)
	}

	return &Service{
		repo: roadmapgrouprepo.New(db),
	}, nil
}
