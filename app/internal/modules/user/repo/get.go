package userrepo

import (
	"context"
	"errors"
	usermodel "roadmap/app/internal/modules/user/model"

	"gorm.io/gorm"
)

func (r *SqliteRepo) FindOne(ctx context.Context, q usermodel.FindQueryBuilder) (usermodel.User, error) {
	whereClause, args, err := q.Build()
	if err != nil {
		return usermodel.User{}, err
	}

	var user usermodel.User
	err = r.db.WithContext(ctx).Where(whereClause, args...).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return usermodel.User{}, usermodel.ErrUserNotFound
		}
		return usermodel.User{}, err
	}

	return user, nil
}
