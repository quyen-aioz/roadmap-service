package roadmapgrouprepo

import (
	"context"
	roadmapgroupmodel "roadmap/app/internal/modules/roadmapgroup/model"
)

func (r *SqliteRepo) GetRoadmapGroup(ctx context.Context) ([]roadmapgroupmodel.RoadmapGroup, error) {
	var roadmapGroups []roadmapgroupmodel.RoadmapGroup
	err := r.db.WithContext(ctx).Raw(`
		SELECT id, name, svg_url, sort_order, is_active, created_at, updated_at, deleted_at
		FROM roadmap_group
		WHERE deleted_at IS NULL
	`).Scan(&roadmapGroups).Error
	return roadmapGroups, err
}
