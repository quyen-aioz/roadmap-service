package authservice

import (
	"context"
	authconfig "roadmap/app/internal/modules/auth/config"
	authmodel "roadmap/app/internal/modules/auth/model"
	usermodel "roadmap/app/internal/modules/user/model"
	"roadmap/pkg/jwtx"
)

func (svc *Service) Login(ctx context.Context, req *authmodel.LoginReq) (*authmodel.LoginResp, error) {
	user, err := svc.userSvc.VerifyUser(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	token, err := svc.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &authmodel.LoginResp{
		AccessToken: token,
	}, nil
}

func (svc *Service) GenerateToken(user usermodel.User) (string, error) {
	claims := jwtx.UserClaim{
		UserID: user.ID,
	}

	return jwtx.GenerateToken(claims, authconfig.AccessTokenDuration)
}

func (svc *Service) Register(ctx context.Context, req *authmodel.RegisterReq) (*authmodel.RegisterResp, error) {
	_, err := svc.userSvc.GetUserByUsername(ctx, req.Username)
	if err == nil {
		return nil, usermodel.ErrUserAlreadyExists
	}

	user, err := svc.userSvc.CreateUser(ctx, usermodel.CreateUserReq{
		Password: req.Password,
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}

	token, err := svc.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &authmodel.RegisterResp{
		AccessToken: token,
	}, nil
}

func (svc *Service) ChangePassword(ctx context.Context, userID string, req *authmodel.ChangePasswordReq) (*authmodel.ChangePasswordResp, error) {
	user, err := svc.userSvc.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if _, err := svc.userSvc.VerifyUser(ctx, user.Username, req.OldPassword); err != nil {
		return nil, usermodel.ErrInvalidPassword
	}

	updatedUser, err := svc.userSvc.UpdatePassword(ctx, user.ID, req.NewPassword)
	if err != nil {
		return nil, err
	}

	token, err := svc.GenerateToken(updatedUser)
	if err != nil {
		return nil, err
	}

	return &authmodel.ChangePasswordResp{
		AccessToken: token,
	}, nil
}
