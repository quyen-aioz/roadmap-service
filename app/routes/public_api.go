package routes

import (
	"fmt"
	authapi "roadmap/app/internal/modules/auth/api"
	roadmapapi "roadmap/app/internal/modules/roadmap/api"
	roadmapgroupapi "roadmap/app/internal/modules/roadmapgroup/api"
	"roadmap/pkg/humax"
)

func registerPublicAPIv1(api humax.API) error {
	roadmapAPI := api.Group("/roadmap", "Roadmap")
	if err := roadmapapi.Init(roadmapAPI); err != nil {
		return fmt.Errorf("register roadmap api: %w", err)
	}

	roadmapGroupAPI := api.Group("/roadmap-group", "Roadmap Group")
	if err := roadmapgroupapi.Init(roadmapGroupAPI); err != nil {
		return fmt.Errorf("register roadmap group api: %w", err)
	}

	authAPI := api.Group("/auth", "Auth")
	if err := authapi.Init(authAPI); err != nil {
		return fmt.Errorf("register auth api: %w", err)
	}

	return nil
}
