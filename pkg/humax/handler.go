package humax

import (
	"context"
	"roadmap/pkg/httpx"

	"github.com/danielgtaylor/huma/v2"
)

type API struct {
	huma.API

	group string
	tags  []string
}

// NewAPI returns a new API which decorates the given API with the given group and tags.
func NewAPI(api huma.API, group string, tags ...string) API {
	return API{
		API:   api,
		group: group,
		tags:  tags,
	}
}

func (api API) Group(path string, tags ...string) API {
	return API{
		API:   api.API,
		group: api.group + path,
		tags:  append(api.tags, tags...),
	}
}

type Operation struct {
	huma.Operation
}

type Output[T any] struct {
	Body httpx.ResponseT[T] `json:",inline" doc:",inline"`
}

func (o Operation) prepare(api API) huma.Operation {
	security := make(map[string][]string)

	operation := o.Operation
	operation.Tags = append(api.tags, operation.Tags...)
	operation.Path = api.group + operation.Path
	operation.Security = append(operation.Security, security)
	// quyen@note: middlewares append later (if need)

	return operation
}

func Register[I, O any](api API, operation Operation, h func(context.Context, *I) (*O, error)) {
	huma.Register(api, operation.prepare(api), func(ctx context.Context, i *I) (*Output[*O], error) {
		output, err := h(ctx, i)
		return &Output[*O]{Body: httpx.AutoResponseT(output, err)}, nil
	})
}

func RegisterWithoutWrapper[I, O any](api API, op Operation, h func(context.Context, *I) (*O, error)) {
	huma.Register(api, op.prepare(api), func(ctx context.Context, i *I) (*O, error) {
		return h(ctx, i)
	})
}

type Empty = struct{}
