package humax

import "github.com/danielgtaylor/huma/v2"

func DefaultConfig() huma.Config {
	config := huma.DefaultConfig("Roadmap Service", "1.0.0")

	return config
}
