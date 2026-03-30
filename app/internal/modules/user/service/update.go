package userservice

import (
	"context"
	usermodel "roadmap/app/internal/modules/user/model"

	"golang.org/x/crypto/bcrypt"
)

func (svc *Service) CreateUser(ctx context.Context, req usermodel.CreateUserReq) (usermodel.User, error) {
	_, err := svc.GetUserByUsername(ctx, req.Username)
	if err == nil {
		return usermodel.User{}, usermodel.ErrUserAlreadyExists
	}

	return svc.repo.Create(ctx, req)
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

func comparePassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
