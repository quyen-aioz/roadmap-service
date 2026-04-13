package roadmaprepo

import (
	"context"
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"
	"time"
)

func (r *SqliteRepo) Create(ctx context.Context, roadmap *roadmapmodel.Roadmap) (string, error) {
	query := `
		INSERT INTO roadmap (id, title, content, status, group_id, cta_label, cta_link, start_date, end_date, thumbnail_url, thumbnail_type, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	err := r.db.WithContext(ctx).Exec(query,
		roadmap.ID,
		roadmap.Title,
		roadmap.Content,
		roadmap.Status,
		roadmap.GroupID,
		roadmap.CTALabel,
		roadmap.CTALink,
		roadmap.StartDate,
		roadmap.EndDate,
		roadmap.ThumbnailURL,
		roadmap.ThumbnailType,
		roadmap.CreatedAt,
		roadmap.UpdatedAt,
	).Error

	if err != nil {
		return "", err
	}

	return roadmap.ID, nil
}

func (r *SqliteRepo) Update(ctx context.Context, id string, u roadmapmodel.RoadmapUpdateBuilder) (string, error) {
	setClause, args, err := u.Build()
	if err != nil {
		return "", err
	}

	baseQuery := "UPDATE roadmap"
	fullQuery := baseQuery + setClause + " WHERE id = ? AND deleted_at IS NULL"
	args = append(args, id)

	err = r.db.WithContext(ctx).Exec(fullQuery, args...).Error
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *SqliteRepo) Delete(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Exec(`
		UPDATE roadmap
		SET deleted_at = ?
		WHERE id = ?
	`, time.Now(), id).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *SqliteRepo) UpdateRoadmapContent(ctx context.Context, req roadmapmodel.UpdateRoadmapContentReq) error {
	updates := map[string]any{}

	if req.Title != nil {
		updates["title"] = *req.Title
	}
	if req.Description != nil {
		updates["description"] = *req.Description
	}
	if req.Content != nil {
		updates["content"] = *req.Content
	}

	if len(updates) == 0 {
		return nil
	}

	return r.db.WithContext(ctx).
		Model(&roadmapmodel.RoadmapContent{}).
		Where("id = ? AND deleted_at IS NULL", roadmapmodel.RoadmapContentID).
		Updates(updates).Error
}
