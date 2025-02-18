package di

import (
	"github.com/reearth/server-scaffold/internal/transport/echo"
	"github.com/reearth/server-scaffold/internal/usecase"
)

func newEchoConfig(conf *Config, usecases usecase.Usecases) echo.Config {
	return echo.Config{
		Dev:      conf.Dev,
		Port:     conf.Port,
		Usecases: usecases,
	}
}
