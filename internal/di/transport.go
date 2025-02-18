package di

import (
	"github.com/reearth/server-scaffold/internal/transport/echo"
	"github.com/reearth/server-scaffold/internal/usecase"
)

func NewEchoConfig(conf *Config, usecases usecase.Usecases) echo.Config {
	return echo.Config{Port: conf.Port, Usecases: usecases, Dev: conf.Dev}
}
