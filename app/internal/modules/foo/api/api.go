package fooapi

import (
	"net/http"
	"roadmap/pkg/humax"

	"github.com/danielgtaylor/huma/v2"
)

type Handler struct {
}

func Init(api humax.API) error {
	handler := &Handler{}

	// GET /v1/foo
	humax.Register(api, humax.Operation{
		Operation: huma.Operation{
			OperationID: "Get Foo",
			Method:      http.MethodGet,
			Path:        "",
		},
	}, handler.GetFoo)

	return nil
}
