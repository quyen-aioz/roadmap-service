package roadmapgroupservice

import (
	"context"
	"fmt"
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"
	roadmapgroupmodel "roadmap/app/internal/modules/roadmapgroup/model"
	"time"
)

func (svc *Service) ReorderGroup(ctx context.Context, orderedIDs []string) ([]roadmapgroupmodel.RoadmapGroup, error) {
	if len(orderedIDs) == 0 {
		return []roadmapgroupmodel.RoadmapGroup{}, nil
	}

	ordered := make([]roadmapgroupmodel.RoadmapGroup, len(orderedIDs))
	for i, id := range orderedIDs {
		ordered[i] = roadmapgroupmodel.RoadmapGroup{
			ID:        roadmapmodel.GroupID(id),
			SortOrder: i,
			UpdatedAt: time.Now(),
		}
	}

	columns := []string{"sort_order", "updated_at"}
	if err := svc.repo.BulkUpsert(ctx, ordered, columns...); err != nil {
		return nil, fmt.Errorf("reorder group: %w", err)
	}

	return ordered, nil
}
