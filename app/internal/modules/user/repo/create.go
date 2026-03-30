package userrepo

import (
	"context"
	usermodel "roadmap/app/internal/modules/user/model"

	"golang.org/x/crypto/bcrypt"
)

func (r *SqliteRepo) Create(ctx context.Context, req usermodel.CreateUserReq) (usermodel.User, error) {
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return usermodel.User{}, err
	}

	user := usermodel.User{
		Username: req.Username,
		Password: hashedPassword,
	}

	err = r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return usermodel.User{}, err
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
