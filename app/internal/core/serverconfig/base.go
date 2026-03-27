package serverconfig

import (
	"roadmap/pkg/envx"
	"roadmap/pkg/srvconf"
	"sync"
	"sync/atomic"
)

func (s ServerAppConfig) GetDir() string {
	return "configs"
}

var (
	instance ServerAppConfig
	loadMu   sync.Mutex
	loaded   atomic.Bool
)

func Init() {
	loadMu.Lock()
	defer loadMu.Unlock()

	srvconf.Load(&instance)
	loaded.Store(true)
}

func Get() ServerAppConfig {
	return instance
}

func _isPROD() bool {
	if !loaded.Load() {
		return false
	}
	return envx.IsPROD(instance.Server.Env)
}

// IsPROD returns true if the server is running in a production environment.
func IsPROD() bool {
	return _isPROD()
}

// IsNonPROD returns true if the server is running in non-production mode.
func IsNonPROD() bool {
	return !IsPROD()
}
