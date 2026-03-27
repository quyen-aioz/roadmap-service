package humax

import (
	"encoding/json"
	"roadmap/pkg/httpx"

	"github.com/danielgtaylor/huma/v2"
)

func ErrorResponse(ctx huma.Context, status int, err error) {
	resp := httpx.ErrorResponseT[any](err)

	ctx.SetHeader("Content-Type", "application/json")
	ctx.SetStatus(status)

	encoder := json.NewEncoder(ctx.BodyWriter())
	if err := encoder.Encode(resp); err != nil {
		panic(err)
	}
}
