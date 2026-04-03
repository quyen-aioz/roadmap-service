package roadmapmodel

import (
	"time"

	"gorm.io/gorm"
)

type Roadmap struct {
	ID        string `gorm:"primaryKey"`
	Title     string
	Content   string
	Status    Status
	GroupID   GroupID
	CTALabel  string
	CTALink   string
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// quyen@note: gorm hook
func (r *Roadmap) BeforeCreate(_ *gorm.DB) (err error) {
	if r.ID == "" {
		r.ID = GenerateHexID()
	}

	return nil
}

func (Roadmap) TableName() string {
	return "roadmap"
}

type UpdateRoadmapReq struct {
	Title     *string
	Content   *string
	Status    *Status
	GroupID   *GroupID
	CTALabel  *string
	CTALink   *string
	StartDate *time.Time
	EndDate   *time.Time
}
