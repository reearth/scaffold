package boot

import "github.com/reearth/server-scaffold/internal/usecase"

func InitPolicies(cfg *Config) usecase.Policies {
	return usecase.DefaultPolicies() // TODO: we can cerbos or other policy engine here
}
