package serverconfig

import (
	"roadmap/pkg/srvconf"
	"sync"
)

func (s ServerAppConfig) GetDir() string {
	return "configs"
}

var (
	instance ServerAppConfig
	loadMu   sync.Mutex
)

func Init() {
	loadMu.Lock()
	defer loadMu.Unlock()

	srvconf.Load(&instance)
}

func Get() ServerAppConfig {
	return instance
}
