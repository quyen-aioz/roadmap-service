package roadmapmodel

import (
	"time"

	"gorm.io/gorm"
)

type Roadmap struct {
	ID            string `gorm:"primaryKey"`
	Title         string
	Content       string
	Status        Status
	GroupID       GroupID
	CTALabel      string
	CTALink       string
	StartDate     time.Time
	EndDate       time.Time
	ThumbnailURL  string
	ThumbnailType string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
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
	Title         *string
	Content       *string
	Status        *Status
	GroupID       *GroupID
	CTALabel      *string
	CTALink       *string
	StartDate     *time.Time
	EndDate       *time.Time
	ThumbnailURL  *string
	ThumbnailType *string
}

type (
	PresignUploadReq struct {
		TaskID      string
		ContentType string
	}

	PresignUploadResp struct {
		UploadURL string
		PublicURL string
	}
)

type RoadmapContent struct {
	ID          string `gorm:"primaryKey"`
	Title       string
	Description string
	Content     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (RoadmapContent) TableName() string {
	return "roadmap_content"
}

type UpdateRoadmapContentReq struct {
	Title       *string
	Description *string
	Content     *string
}
