package serverconfig

import "roadmap/pkg/srvconf"

var _ srvconf.Configuration = (*ServerAppConfig)(nil)

type ServerAppConfig struct {
	srvconf.BaseConfig `mapstructure:",squash"`
}
