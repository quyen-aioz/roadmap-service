package authapi

import (
	"net/http"
	authservice "roadmap/app/internal/modules/auth/service"
	"roadmap/pkg/humax"

	"github.com/danielgtaylor/huma/v2"
)

type Handler struct {
	svc *authservice.Service
}

func New() (*Handler, error) {
	svc, err := authservice.New()
	if err != nil {
		return nil, err
	}

	return &Handler{
		svc: svc,
	}, nil
}

func Init(api humax.API) error {
	handler, err := New()
	if err != nil {
		return err
	}

	middlewares := []humax.Middleware{
		humax.RequireAccessToken(),
	}

	// POST /v1/auth/login
	humax.Register(api, humax.Operation{
		Operation: huma.Operation{
			OperationID: "Login",
			Method:      http.MethodPost,
			Path:        "/login",
		},
	}, handler.Login)

	// GET /v1/auth/me
	humax.Register(api, humax.Operation{
		Operation: huma.Operation{
			OperationID: "Me",
			Method:      http.MethodGet,
			Path:        "/me",
		},
		CustomMiddlewares: middlewares,
	}, handler.GetMe)

	// quyen@note: not allow to register, just call once to create admin user
	// // POST /v1/auth/register
	// humax.Register(api, humax.Operation{
	// 	Operation: huma.Operation{
	// 		OperationID: "Register",
	// 		Method:      http.MethodPost,
	// 		Path:        "/register",
	// 	},
	// }, handler.Register)

	return nil
}
