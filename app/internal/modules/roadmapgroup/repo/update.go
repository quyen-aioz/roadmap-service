package roadmapgrouprepo

import (
	"context"
	roadmapgroupmodel "roadmap/app/internal/modules/roadmapgroup/model"

	"gorm.io/gorm/clause"
)

func (r *SqliteRepo) BulkUpsert(ctx context.Context, groups []roadmapgroupmodel.RoadmapGroup, updateColumns ...string) error {
	if len(groups) == 0 {
		return nil
	}

	tx := r.db.WithContext(ctx)

	if len(updateColumns) > 0 {
		tx = tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns(updateColumns),
		})
	} else {
		tx = tx.Clauses(clause.OnConflict{UpdateAll: true})
	}

	return tx.Create(&groups).Error
}
