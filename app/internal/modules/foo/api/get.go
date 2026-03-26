package fooapi

import (
	"context"
	"roadmap/pkg/humax"
)

type (
	GetFooResponse struct {
		Name string `json:"name"`
	}
)

func (h *Handler) GetFoo(ctx context.Context, _ *humax.Empty) (*GetFooResponse, error) {
	return &GetFooResponse{
		Name: "wakanda",
	}, nil
}
