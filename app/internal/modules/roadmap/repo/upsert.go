package roadmaprepo

import (
	"context"
	"database/sql"
	"fmt"
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"
	"strings"
	"time"
)

type dbExecutor interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

func (r *SqliteRepo) BulkUpsert(ctx context.Context, roadmaps []roadmapmodel.Roadmap) error {
	return r.bulkUpsert(ctx, r.db, roadmaps)
}

func (r *SqliteRepo) BulkDelete(ctx context.Context, deleteIDs []string) error {
	return r.bulkDelete(ctx, r.db, deleteIDs)
}

func (r *SqliteRepo) Sync(ctx context.Context, roadmaps []roadmapmodel.Roadmap, deleteIDs []string) ([]roadmapmodel.Roadmap, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	if err := r.bulkUpsert(ctx, tx, roadmaps); err != nil {
		return nil, err
	}

	if err := r.bulkDelete(ctx, tx, deleteIDs); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return r.GetRoadmap(ctx)
}

func (r *SqliteRepo) bulkUpsert(ctx context.Context, exec dbExecutor, roadmaps []roadmapmodel.Roadmap) error {
	if len(roadmaps) == 0 {
		return nil
	}

	query := `
		INSERT INTO roadmap (id, title, content, status, group_id, start_date, end_date, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			title = excluded.title,
			content = excluded.content,
			status = excluded.status,
			group_id = excluded.group_id,
			start_date = excluded.start_date,
			end_date = excluded.end_date,
			updated_at = excluded.updated_at
	`

	now := time.Now()

	for _, roadmap := range roadmaps {
		id := roadmap.ID
		if id == "" {
			id = GenerateHexID()
		}

		createdAt := roadmap.CreatedAt
		if createdAt.IsZero() {
			createdAt = now
		}

		_, err := exec.ExecContext(ctx, query,
			id,
			roadmap.Title,
			roadmap.Content,
			roadmap.Status,
			roadmap.GroupID,
			roadmap.StartDate,
			roadmap.EndDate,
			createdAt,
			now,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *SqliteRepo) bulkDelete(ctx context.Context, exec dbExecutor, deleteIDs []string) error {
	if len(deleteIDs) == 0 {
		return nil
	}

	placeholders := make([]string, len(deleteIDs))
	args := make([]any, len(deleteIDs)+1)
	args[0] = time.Now()
	for i, id := range deleteIDs {
		placeholders[i] = "?"
		args[i+1] = id
	}

	deleteQuery := fmt.Sprintf(`
		UPDATE roadmap
		SET deleted_at = ?
		WHERE id IN (%s) AND deleted_at IS NULL
	`, strings.Join(placeholders, ", "))

	_, err := exec.ExecContext(ctx, deleteQuery, args...)
	return err
}
