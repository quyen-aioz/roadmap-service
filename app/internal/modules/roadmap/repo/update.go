package roadmaprepo

import (
	"context"
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"
	"time"
)

func (r *SqliteRepo) Create(ctx context.Context, roadmap *roadmapmodel.Roadmap) (string, error) {
	newID := GenerateHexID()
	query := `
		INSERT INTO roadmap (id, title, content, status, group_id, start_date, end_date, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.ExecContext(ctx, query,
		newID,
		roadmap.Title,
		roadmap.Content,
		roadmap.Status,
		roadmap.GroupID,
		roadmap.StartDate,
		roadmap.EndDate,
		roadmap.CreatedAt,
		roadmap.UpdatedAt,
	)

	if err != nil {
		return "", err
	}

	return newID, nil
}

func (r *SqliteRepo) Update(ctx context.Context, id string, u roadmapmodel.RoadmapUpdateBuilder) (string, error) {
	setClause, args, err := u.Build()
	if err != nil {
		return "", err
	}

	baseQuery := "UPDATE roadmap"
	fullQuery := baseQuery + setClause + " WHERE id = ? AND deleted_at IS NULL"
	args = append(args, id)

	_, err = r.db.ExecContext(ctx, fullQuery, args...)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *SqliteRepo) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE roadmap
		SET deleted_at = ?
		WHERE id = ?
	`, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}
