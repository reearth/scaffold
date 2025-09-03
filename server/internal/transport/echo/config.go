package echo

import (
	"github.com/reearth/reearthx/appx"
	"github.com/reearth/scaffold/server/internal/usecase"
)

type Config struct {
	Port         string
	Usecases     usecase.Usecases
	Dev          bool
	JWTProviders []appx.JWTProvider
	HealthCheck  HealthCheckConfig
	Version      string
	DB           string
}

type HealthCheckConfig struct {
	Username string
	Password string
}
