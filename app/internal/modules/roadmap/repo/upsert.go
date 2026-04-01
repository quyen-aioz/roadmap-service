package roadmaprepo

import (
	"context"
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *SqliteRepo) BulkUpsert(ctx context.Context, roadmaps []roadmapmodel.Roadmap) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return r.bulkUpsert(ctx, tx, roadmaps)
	})
}

func (r *SqliteRepo) BulkDelete(ctx context.Context, deleteIDs []string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return r.bulkDelete(ctx, tx, deleteIDs)
	})
}

func (r *SqliteRepo) Sync(ctx context.Context, roadmaps []roadmapmodel.Roadmap, deleteIDs []string) ([]roadmapmodel.Roadmap, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := r.bulkUpsert(ctx, tx, roadmaps); err != nil {
			return err
		}

		if err := r.bulkDelete(ctx, tx, deleteIDs); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return r.GetRoadmap(ctx)
}

// func (r *SqliteRepo) bulkUpsert(ctx context.Context, db *gorm.DB, roadmaps []roadmapmodel.Roadmap) error {
// 	if len(roadmaps) == 0 {
// 		return nil
// 	}

// 	query := `
// 		INSERT INTO roadmap (id, title, content, status, group_id, start_date, end_date, created_at, updated_at)
// 		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
// 		ON CONFLICT(id) DO UPDATE SET
// 			title = excluded.title,
// 			content = excluded.content,
// 			status = excluded.status,
// 			group_id = excluded.group_id,
// 			start_date = excluded.start_date,
// 			end_date = excluded.end_date,
// 			updated_at = excluded.updated_at
// 	`

// 	now := time.Now()

// 	for _, roadmap := range roadmaps {
// 		id := roadmap.ID
// 		if id == "" {
// 			id = GenerateHexID()
// 		}

// 		createdAt := roadmap.CreatedAt
// 		if createdAt.IsZero() {
// 			createdAt = now
// 		}

// 		err := db.WithContext(ctx).Exec(query,
// 			id,
// 			roadmap.Title,
// 			roadmap.Content,
// 			roadmap.Status,
// 			roadmap.GroupID,
// 			roadmap.StartDate,
// 			roadmap.EndDate,
// 			createdAt,
// 			now,
// 		).Error
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (r *SqliteRepo) bulkDelete(ctx context.Context, db *gorm.DB, deleteIDs []string) error {
// 	if len(deleteIDs) == 0 {
// 		return nil
// 	}

// 	placeholders := make([]string, len(deleteIDs))
// 	args := make([]any, len(deleteIDs)+1)
// 	args[0] = time.Now()
// 	for i, id := range deleteIDs {
// 		placeholders[i] = "?"
// 		args[i+1] = id
// 	}

// 	deleteQuery := fmt.Sprintf(`
// 		UPDATE roadmap
// 		SET deleted_at = ?
// 		WHERE id IN (%s) AND deleted_at IS NULL
// 	`, strings.Join(placeholders, ", "))

// 	return db.WithContext(ctx).Exec(deleteQuery, args...).Error
// }

func (r *SqliteRepo) bulkUpsert(ctx context.Context, db *gorm.DB, roadmaps []roadmapmodel.Roadmap) error {
	if len(roadmaps) == 0 {
		return nil
	}

	// GORM handles the "ON CONFLICT" and "updated_at" automatically.
	// clause.OnConflict is the standard way to handle Upserts in GORM.
	return db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true, // This updates all fields with values from the struct
	}).Create(&roadmaps).Error
}

func (r *SqliteRepo) bulkDelete(ctx context.Context, db *gorm.DB, deleteIDs []string) error {
	if len(deleteIDs) == 0 {
		return nil
	}

	// Since your model has gorm.DeletedAt, this automatically:
	// 1. Checks if it's already null
	// 2. Performs an UPDATE with the current timestamp
	// 3. Handles the "IN" clause for the slice of IDs
	return db.WithContext(ctx).Delete(&roadmapmodel.Roadmap{}, deleteIDs).Error
}
