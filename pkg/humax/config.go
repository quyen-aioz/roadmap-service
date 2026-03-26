package humax

import "github.com/danielgtaylor/huma/v2"

func DefaultConfig() huma.Config {
	config := huma.DefaultConfig("Roadmap Service", "")

	return config
}
