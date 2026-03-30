package humax

import "github.com/danielgtaylor/huma/v2"

func DefaultConfig() huma.Config {
	config := huma.DefaultConfig("Roadmap Service", "1.0.0")
	config.OpenAPI.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"bearer": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}

	return config
}
