package roadmapmodel

import "time"

type (
	RoadmapDTO struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		Status    Status    `json:"status"`
		GroupID   GroupID   `json:"group_id"`
		CTALabel  string    `json:"cta_label"`
		CTALink   string    `json:"cta_link"`
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

type (
	RoadmapContentDTO struct {
		ID          string    `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Content     string    `json:"content"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
