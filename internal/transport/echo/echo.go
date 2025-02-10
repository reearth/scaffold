package echo

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/reearth/server-scaffold/internal/boot"
	"github.com/reearth/server-scaffold/internal/transport/gql"
	"github.com/reearth/server-scaffold/internal/usecase"
	"golang.org/x/net/http2"
)

type Server struct {
	echo     *echo.Echo
	config   Config
	http2srv *http2.Server
}

type Config struct {
	Port     string
	Usecases usecase.Usecases
	Dev      bool
}

func NewEchoConfig(conf *boot.Config, usecases usecase.Usecases, dev bool) Config {
	return Config{Port: conf.Port, Usecases: usecases, Dev: dev}
}

func New(conf Config) *Server {
	e := echo.New()
	if err := initEcho(e, conf); err != nil {
		panic(err)
	}

	s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}

	return &Server{echo: e, config: conf, http2srv: s}
}

func (s *Server) Start() error {
	s.echo.Logger.Infof("Starting server on :%s", s.config.Port)
	return s.echo.StartH2CServer(fmt.Sprintf(":%s", s.config.Port), s.http2srv)
}

func initEcho(e *echo.Echo, conf Config) error {
	e.HideBanner = true

	e.Use(
		middleware.Recover(),
		middleware.RequestID(),
		middleware.Logger(),
	)

	api := e.Group("/api")

	api.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	api.POST("/graphql", echo.WrapHandler(gql.NewServer(conf.Usecases)))
	if conf.Dev {
		api.GET("/graphql", echo.WrapHandler(gql.Playground("/graphql")))
	}

	return nil
}
