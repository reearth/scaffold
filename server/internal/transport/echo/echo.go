package echo

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	glog "github.com/labstack/gommon/log"
	"github.com/reearth/reearthx/appx"
	"github.com/reearth/reearthx/log"
	"github.com/reearth/scaffold/server/internal/transport/gql"
	"golang.org/x/net/http2"
)

type Server struct {
	echo     *echo.Echo
	config   Config
	http2srv *http2.Server
}

func New(conf Config) *Server {
	e := echo.New()
	if err := initEcho(e, conf); err != nil {
		panic(err)
	}

	s := &http2.Server{}
	return &Server{echo: e, config: conf, http2srv: s}
}

func (s *Server) Start() error {
	s.echo.Logger.Infof("Starting server on :%s", s.config.Port)
	return s.echo.StartH2CServer(fmt.Sprintf(":%s", s.config.Port), s.http2srv)
}

func initEcho(e *echo.Echo, conf Config) error {
	e.HideBanner = true
	logger := log.NewEcho()
	e.Logger = logger

	// middleware
	jwtMiddleware, err := appx.AuthMiddleware(conf.JWTProviders, nil, true)
	if err != nil {
		return fmt.Errorf("failed to create auth middleware: %w", err)
	}

	e.Use(
		middleware.RecoverWithConfig(middleware.RecoverConfig{
			LogLevel: glog.ERROR,
		}),
		middleware.Gzip(),
		echo.WrapMiddleware(appx.RequestIDMiddleware()),
		logger.AccessLogger(),
		echo.WrapMiddleware(jwtMiddleware),
	)

	// routes
	api := e.Group("/api")

	api.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	// health check endpoint
	api.GET("/health", HealthCheck(conf))

	api.POST("/graphql", echo.WrapHandler(gql.NewServer(conf.Usecases, conf.Dev)))
	if conf.Dev {
		api.GET("/graphql", echo.WrapHandler(gql.Playground("/graphql")))
	}

	return nil
}
