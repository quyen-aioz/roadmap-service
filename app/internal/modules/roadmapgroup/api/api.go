package roadmapgroupapi

import (
	"net/http"
	roadmapgroupservice "roadmap/app/internal/modules/roadmapgroup/service"
	"roadmap/pkg/humax"

	"github.com/danielgtaylor/huma/v2"
)

type Handler struct {
	svc *roadmapgroupservice.Service
}

func New() (*Handler, error) {
	svc, err := roadmapgroupservice.New()
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
	// GET /v1/roadmap-group
	humax.Register(api, humax.Operation{
		Operation: huma.Operation{
			OperationID: "Get roadmap group",
			Method:      http.MethodGet,
			Path:        "",
		},
	}, handler.GetRoadmapGroup)

	// POST /v1/roadmap-group/reorder
	humax.Register(api, humax.Operation{
		Operation: huma.Operation{
			OperationID: "Reorder roadmap group",
			Method:      http.MethodPost,
			Path:        "/reorder",
		},
		CustomMiddlewares: middlewares,
	}, handler.ReorderGroup)

	return nil
}
