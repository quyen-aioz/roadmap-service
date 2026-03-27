package roadmapservice

import (
	"context"
	"fmt"
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"
)

func (svc *Service) CreateRoadmap(ctx context.Context, req *roadmapmodel.Roadmap) (string, error) {
	id, err := svc.repo.Create(ctx, req)
	if err != nil {
		return "", fmt.Errorf("create roadmap: %w", err)
	}

	return id, nil
}

func (svc *Service) UpdateRoadmap(ctx context.Context, id string, req roadmapmodel.UpdateRoadmapReq) (roadmapmodel.Roadmap, error) {
	_, err := svc.repo.FindOne(ctx, roadmapmodel.FindQueryBuilder{
		ID: id,
	})

	if err != nil {
		return roadmapmodel.Roadmap{}, fmt.Errorf("road map with ID: %s not found: %w", id, err)
	}

	builder := roadmapmodel.RoadmapUpdateBuilder(req)

	_, err = svc.repo.Update(ctx, id, builder)
	if err != nil {
		return roadmapmodel.Roadmap{}, fmt.Errorf("update roadmap: %w", err)
	}

	return svc.repo.FindOne(ctx, roadmapmodel.FindQueryBuilder{
		ID: id,
	})
}

func (svc *Service) DeleteRoadmap(ctx context.Context, id string) error {
	_, err := svc.repo.FindOne(ctx, roadmapmodel.FindQueryBuilder{
		ID: id,
	})
	if err != nil {
		return fmt.Errorf("road map with ID: %s not found: %w", id, err)
	}

	err = svc.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("delete roadmap: %w", err)
	}

	return nil
}
