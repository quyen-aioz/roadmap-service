package authapi

import (
	"context"
	"roadmap/pkg/humax"
)

type (
	MeRespond struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	}
)

func (h *Handler) GetMe(ctx context.Context, _ *humax.Empty) (*MeRespond, error) {
	userID, err := humax.GetUserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	user, err := h.svc.GetMe(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &MeRespond{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}
