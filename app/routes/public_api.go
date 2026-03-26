package routes

import (
	"fmt"
	fooapi "roadmap/app/internal/modules/foo/api"
	"roadmap/pkg/humax"
)

func registerPublicAPIv1(api humax.API) error {
	fooAPI := api.Group("/foo", "Foo")
	if err := fooapi.Init(fooAPI); err != nil {
		return fmt.Errorf("register foo api: %w", err)
	}

	return nil
}
