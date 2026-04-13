package roadmaprepo

import (
	"context"
	"errors"
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"

	"gorm.io/gorm"
)

func (r *SqliteRepo) GetRoadmap(ctx context.Context) ([]roadmapmodel.Roadmap, error) {
	var roadmaps []roadmapmodel.Roadmap
	err := r.db.WithContext(ctx).Raw(`
        SELECT id, title, content, status, group_id, cta_label, cta_link, start_date, end_date, thumbnail_url, thumbnail_type, created_at, updated_at, deleted_at
        FROM roadmap
        WHERE deleted_at IS NULL
    `).Scan(&roadmaps).Error
	return roadmaps, err
}

func (r *SqliteRepo) GetRoadmapContent(ctx context.Context) (roadmapmodel.RoadmapContent, error) {
	var roadmapContent roadmapmodel.RoadmapContent
	err := r.db.WithContext(ctx).Raw(`
        SELECT id, title, description, content, created_at, updated_at, deleted_at
        FROM roadmap_content
        WHERE deleted_at IS NULL
    `).Scan(&roadmapContent).Error
	return roadmapContent, err
}

func (r *SqliteRepo) FindOne(ctx context.Context, q roadmapmodel.FindQueryBuilder) (roadmapmodel.Roadmap, error) {
	whereClause, args := q.Build()

	baseQuery := `SELECT id, title, content, status, group_id, cta_label, cta_link, start_date, end_date, thumbnail_url, thumbnail_type, created_at, updated_at, deleted_at FROM roadmap`
	fullQuery := baseQuery + whereClause + " LIMIT 1"

	var roadmap roadmapmodel.Roadmap
	err := r.db.WithContext(ctx).Raw(fullQuery, args...).Scan(&roadmap).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return roadmapmodel.Roadmap{}, roadmapmodel.ErrRoadmapNotFound
		}
		return roadmapmodel.Roadmap{}, err
	}

	return roadmap, nil
}

func (r *SqliteRepo) GetAllThumbnailURLs(ctx context.Context) ([]string, error) {
	var urls []string
	err := r.db.WithContext(ctx).Raw(`
		SELECT thumbnail_url
		FROM roadmap
		WHERE deleted_at IS NULL
		  AND thumbnail_url != ''
	`).Scan(&urls).Error
	return urls, err
}
