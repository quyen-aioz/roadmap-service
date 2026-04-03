package userservice

import (
	"context"
	"fmt"
	usermodel "roadmap/app/internal/modules/user/model"
)

func (svc *Service) GetUserByID(ctx context.Context, id string) (usermodel.User, error) {
	user, err := svc.repo.FindOne(ctx, usermodel.FindQueryBuilder{
		ID: id,
	})
	if err != nil {
		return usermodel.User{}, fmt.Errorf("get user by id: %w", err)
	}

	return user, nil
}

func (svc *Service) GetUserByUsername(ctx context.Context, username string) (usermodel.User, error) {
	user, err := svc.repo.FindOne(ctx, usermodel.FindQueryBuilder{
		Username: username,
	})
	if err != nil {
		return usermodel.User{}, fmt.Errorf("get user by username: %w", err)
	}

	return user, nil
}
