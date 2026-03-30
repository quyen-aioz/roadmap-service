package authservice

import (
	userservice "roadmap/app/internal/modules/user/service"
)

type Service struct {
	userSvc *userservice.Service
}

func New() (*Service, error) {
	userSvc, err := userservice.New()
	if err != nil {
		return nil, err
	}

	return &Service{
		userSvc: userSvc,
	}, nil
}
