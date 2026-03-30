package routes

import (
	"fmt"
	authapi "roadmap/app/internal/modules/auth/api"
	fooapi "roadmap/app/internal/modules/foo/api"
	roadmapapi "roadmap/app/internal/modules/roadmap/api"
	"roadmap/pkg/humax"
)

func registerPublicAPIv1(api humax.API) error {
	fooAPI := api.Group("/foo", "Foo")
	if err := fooapi.Init(fooAPI); err != nil {
		return fmt.Errorf("register foo api: %w", err)
	}

	roadmapAPI := api.Group("/roadmap", "Roadmap")
	if err := roadmapapi.Init(roadmapAPI); err != nil {
		return fmt.Errorf("register roadmap api: %w", err)
	}

	authAPI := api.Group("/auth", "Auth")
	if err := authapi.Init(authAPI); err != nil {
		return fmt.Errorf("register auth api: %w", err)
	}

	return nil
}
