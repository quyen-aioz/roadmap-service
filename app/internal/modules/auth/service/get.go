package authservice

import (
	"context"
	usermodel "roadmap/app/internal/modules/user/model"
)

func (svc *Service) GetMe(ctx context.Context, userID string) (usermodel.User, error) {
	return svc.userSvc.GetUserByID(ctx, userID)
}
