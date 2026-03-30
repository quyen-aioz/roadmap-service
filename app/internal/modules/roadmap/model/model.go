package roadmapmodel

import "time"

type Roadmap struct {
	ID        string
	Title     string
	Content   string
	Status    Status
	GroupID   GroupID
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type UpdateRoadmapReq struct {
	Title     *string
	Content   *string
	Status    *Status
	GroupID   *GroupID
	StartDate *time.Time
	EndDate   *time.Time
}
