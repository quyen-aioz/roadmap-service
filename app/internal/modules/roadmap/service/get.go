package roadmapservice

import (
	"context"
	"fmt"
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"
)

func (svc *Service) GetRoadmap(ctx context.Context) ([]roadmapmodel.Roadmap, error) {
	roadmap, err := svc.repo.GetRoadmap(ctx)
	if err != nil {
		return nil, fmt.Errorf("get roadmap: %w", err)
	}
	return roadmap, nil
}
