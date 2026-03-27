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

	// GET /v1/roadmap
	humax.Register(api, humax.Operation{
		Operation: huma.Operation{
			OperationID: "Get Roadmap",
			Method:      http.MethodGet,
			Path:        "",
		},
	}, handler.GetRoadmap)

	return nil
}
