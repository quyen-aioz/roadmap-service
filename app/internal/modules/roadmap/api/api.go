package roadmapapi

import (
	"net/http"
	roadmapservice "roadmap/app/internal/modules/roadmap/service"
	"roadmap/pkg/humax"

	"github.com/danielgtaylor/huma/v2"
)

type Handler struct {
	svc *roadmapservice.Service
}

func New() (*Handler, error) {
	svc, err := roadmapservice.New()
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
	// GET /v1/roadmap
	humax.Register(api, humax.Operation{
		Operation: huma.Operation{
			OperationID: "Get Roadmap",
			Method:      http.MethodGet,
			Path:        "",
		},
	}, handler.GetRoadmap)

	// POST /v1/roadmap/sync
	humax.Register(api, humax.Operation{
		Operation: huma.Operation{
			OperationID: "Sync Roadmap",
			Method:      http.MethodPost,
			Path:        "/sync",
		},
		CustomMiddlewares: middlewares,
	}, handler.SyncRoadmap)

	// GET /v1/roadmap/content
	humax.Register(api, humax.Operation{
		Operation: huma.Operation{
			OperationID: "Get roadmap content",
			Method:      http.MethodGet,
			Path:        "/content",
		},
	}, handler.GetRoadmapContent)

	// POST /v1/roadmap/content
	humax.Register(api, humax.Operation{
		Operation: huma.Operation{
			OperationID: "Update roadmap content",
			Method:      http.MethodPost,
			Path:        "/content",
		},
		CustomMiddlewares: middlewares,
	}, handler.UpdateRoadmapContent)

	// // POST /v1/roadmap
	// humax.Register(api, humax.Operation{
	// 	Operation: huma.Operation{
	// 		OperationID: "Create Roadmap",
	// 		Method:      http.MethodPost,
	// 		Path:        "",
	// 	},
	// 	CustomMiddlewares: middlewares,
	// }, handler.CreateRoadmap)

	// // PUT /v1/roadmap/{id}
	// humax.Register(api, humax.Operation{
	// 	Operation: huma.Operation{
	// 		OperationID: "Update Roadmap",
	// 		Method:      http.MethodPut,
	// 		Path:        "/{id}",
	// 	},
	// 	CustomMiddlewares: middlewares,
	// }, handler.UpdateRoadmap)

	// // DELETE /v1/roadmap/{id}
	// humax.Register(api, humax.Operation{
	// 	Operation: huma.Operation{
	// 		OperationID: "Delete Roadmap",
	// 		Method:      http.MethodDelete,
	// 		Path:        "/{id}",
	// 	},
	// 	CustomMiddlewares: middlewares,
	// }, handler.DeleteRoadmap)

	return nil
}
