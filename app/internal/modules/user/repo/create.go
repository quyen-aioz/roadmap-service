package userrepo

import (
	"context"
	usermodel "roadmap/app/internal/modules/user/model"
)

func (r *SqliteRepo) Create(ctx context.Context, req usermodel.CreateUserReq) (usermodel.User, error) {
	user := usermodel.User{
		Username: req.Username,
		Password: req.Password,
	}

	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return usermodel.User{}, err
	}

	return user, nil
}

func (r *SqliteRepo) Update(ctx context.Context, userID string, req usermodel.UpdateUserReq) (usermodel.User, error) {
	user, err := r.FindOne(ctx, usermodel.FindQueryBuilder{
		ID: userID,
	})
	if err != nil {
		return usermodel.User{}, err
	}

	if req.Username != nil {
		user.Username = *req.Username
	}

	if req.Password != nil {
		user.Password = *req.Password
	}

	err = r.db.WithContext(ctx).Where("id = ?", userID).Updates(&user).Error
	if err != nil {
		return usermodel.User{}, err
	}

	return user, nil
}
