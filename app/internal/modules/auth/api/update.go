package authapi

import (
	"context"
	authmodel "roadmap/app/internal/modules/auth/model"
	"roadmap/pkg/humax"
)

type (
	LoginRequest struct {
		Body struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
	}
	LoginRespond struct {
		AccessToken string `json:"access_token"`
	}
)

func (h *Handler) Login(ctx context.Context, req *LoginRequest) (*LoginRespond, error) {
	resp, err := h.svc.Login(ctx, &authmodel.LoginReq{
		Username: req.Body.Username,
		Password: req.Body.Password,
	})
	if err != nil {
		return nil, err
	}

	return &LoginRespond{
		AccessToken: resp.AccessToken,
	}, nil
}

type (
	RegisterRequest struct {
		Body struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
	}
)

func (h *Handler) Register(ctx context.Context, req *RegisterRequest) (*humax.Empty, error) {
	_, err := h.svc.Register(ctx, &authmodel.RegisterReq{
		Username: req.Body.Username,
		Password: req.Body.Password,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
