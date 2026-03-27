package envx

import (
	"strings"
)

type Env = string

const (
	EnvLocal      = "local"
	EnvStaging    = "staging"
	EnvProduction = "production"
)

func IsPROD(env Env) bool {
	return strings.EqualFold(env, EnvProduction)
}

func IsLocal(env Env) bool {
	return strings.EqualFold(env, EnvLocal)
}

func IsStaging(env Env) bool {
	return strings.EqualFold(env, EnvStaging)
}
