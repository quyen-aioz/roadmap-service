package authservice

import (
	"context"
	authconfig "roadmap/app/internal/modules/auth/config"
	authmodel "roadmap/app/internal/modules/auth/model"
	usermodel "roadmap/app/internal/modules/user/model"
	"roadmap/pkg/jwtx"
)

func (s *Service) Login(ctx context.Context, req *authmodel.LoginReq) (*authmodel.LoginResp, error) {
	user, err := s.userSvc.VerifyUser(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	token, err := s.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &authmodel.LoginResp{
		AccessToken: token,
	}, nil
}

func (s *Service) GenerateToken(user usermodel.User) (string, error) {
	claims := jwtx.UserClaim{
		UserID: user.ID,
	}

	return jwtx.GenerateToken(claims, authconfig.AccessTokenDuration)
}

func (s *Service) Register(ctx context.Context, req *authmodel.RegisterReq) (*authmodel.RegisterResp, error) {
	_, err := s.userSvc.GetUserByUsername(ctx, req.Username)
	if err == nil {
		return nil, usermodel.ErrUserAlreadyExists
	}

	user, err := s.userSvc.CreateUser(ctx, usermodel.CreateUserReq{
		Password: req.Password,
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}

	token, err := s.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &authmodel.RegisterResp{
		AccessToken: token,
	}, nil
}
