package roadmaprepo

import (
	"context"
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"
)

func (r *SqliteRepo) GetRoadmap(ctx context.Context) ([]roadmapmodel.Roadmap, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, description, created_at, updated_at, deleted_at
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
			&roadmap.Name,
			&roadmap.Description,
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
