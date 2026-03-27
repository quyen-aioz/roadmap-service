package roadmaprepo

import (
	"context"
	"database/sql"
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"
)

func (r *SqliteRepo) GetRoadmap(ctx context.Context) ([]roadmapmodel.Roadmap, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, title, content, status, group_id, start_date, end_date, created_at, updated_at, deleted_at
		FROM roadmap
		WHERE deleted_at IS NULL
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roadmaps []roadmapmodel.Roadmap

	for rows.Next() {
		var roadmap roadmapmodel.Roadmap
		if err := rows.Scan(
			&roadmap.ID,
			&roadmap.Title,
			&roadmap.Content,
			&roadmap.Status,
			&roadmap.GroupID,
			&roadmap.StartDate,
			&roadmap.EndDate,
			&roadmap.CreatedAt,
			&roadmap.UpdatedAt,
			&roadmap.DeletedAt,
		); err != nil {
			return nil, err
		}
		roadmaps = append(roadmaps, roadmap)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roadmaps, nil
}
func (r *SqliteRepo) FindOne(ctx context.Context, q roadmapmodel.FindQueryBuilder) (roadmapmodel.Roadmap, error) {
	whereClause, args := q.Build()

	baseQuery := `SELECT id, title, content, status, group_id, start_date, end_date, created_at, updated_at, deleted_at FROM roadmap`
	fullQuery := baseQuery + whereClause + " LIMIT 1"

	var roadmap roadmapmodel.Roadmap
	err := r.db.QueryRowContext(ctx, fullQuery, args...).Scan(
		&roadmap.ID,
		&roadmap.Title,
		&roadmap.Content,
		&roadmap.Status,
		&roadmap.GroupID,
		&roadmap.StartDate,
		&roadmap.EndDate,
		&roadmap.CreatedAt,
		&roadmap.UpdatedAt,
		&roadmap.DeletedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return roadmapmodel.Roadmap{}, roadmapmodel.ErrRoadmapNotFound
		}
		return roadmapmodel.Roadmap{}, err
	}

	return roadmap, nil
}
