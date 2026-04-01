package userservice

import (
	"context"
	"fmt"
	usermodel "roadmap/app/internal/modules/user/model"

	"golang.org/x/crypto/bcrypt"
)

func (svc *Service) CreateUser(ctx context.Context, req usermodel.CreateUserReq) (usermodel.User, error) {
	_, err := svc.GetUserByUsername(ctx, req.Username)
	if err == nil {
		return usermodel.User{}, usermodel.ErrUserAlreadyExists
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return usermodel.User{}, fmt.Errorf("hash password: %w", err)
	}
	return svc.repo.Create(ctx, usermodel.CreateUserReq{
		Username: req.Username,
		Password: hashedPassword,
	})
}

func (svc *Service) VerifyUser(ctx context.Context, username, password string) (usermodel.User, error) {
	user, err := svc.GetUserByUsername(ctx, username)
	if err != nil {
		return usermodel.User{}, err
	}

	if !comparePassword(user.Password, password) {
		return usermodel.User{}, usermodel.ErrInvalidPassword
	}

	return user, nil
}

func (svc *Service) UpdatePassword(ctx context.Context, userID, password string) (usermodel.User, error) {
	user, err := svc.GetUserByID(ctx, userID)
	if err != nil {
		return usermodel.User{}, err
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return usermodel.User{}, fmt.Errorf("hash password: %w", err)
	}

	return svc.repo.Update(ctx, user.ID, usermodel.UpdateUserReq{
		Password: &hashedPassword,
	})
}

func comparePassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
