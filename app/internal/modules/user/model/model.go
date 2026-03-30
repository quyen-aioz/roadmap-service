package usermodel

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"primaryKey"`
	Username string
	Password string
}

func (User) TableName() string {
	return "users"
}

// quyen@note: gorm hook
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}

	return nil
}

type (
	CreateUserReq struct {
		Username string
		Password string
	}

	CreateUserResp struct {
		ID       string
		Username string
	}
)
