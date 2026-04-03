package roadmapgroupservice

import (
	"context"
	"fmt"
	roadmapgroupmodel "roadmap/app/internal/modules/roadmapgroup/model"
)

func (svc *Service) GetRoadmapGroup(ctx context.Context) ([]roadmapgroupmodel.RoadmapGroup, error) {
	roadmapGroup, err := svc.repo.GetRoadmapGroup(ctx)
	if err != nil {
		return nil, fmt.Errorf("get roadmap group: %w", err)
	}
	return roadmapGroup, nil
}
