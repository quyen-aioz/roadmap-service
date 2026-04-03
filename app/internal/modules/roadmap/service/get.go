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

func (svc *Service) GetRoadmapByID(ctx context.Context, id string) (roadmapmodel.Roadmap, error) {
	roadmap, err := svc.repo.FindOne(ctx, roadmapmodel.FindQueryBuilder{
		ID: id,
	})
	if err != nil {
		return roadmapmodel.Roadmap{}, fmt.Errorf("get roadmap by id: %w", err)
	}

	return roadmap, nil
}

func (svc *Service) GetRoadmapContent(ctx context.Context) (roadmapmodel.RoadmapContent, error) {
	content, err := svc.repo.GetRoadmapContent(ctx)
	if err != nil {
		return roadmapmodel.RoadmapContent{}, fmt.Errorf("get roadmap content: %w", err)
	}

	return content, nil
}
