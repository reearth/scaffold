package echo

import (
	"log"
	"net/http"
	"time"

	"github.com/hellofresh/health-go/v5"
	"github.com/hellofresh/health-go/v5/checks/mongo"
	"github.com/labstack/echo/v4"
)

// HealthCheck creates a health check handler function
func HealthCheck(conf Config) echo.HandlerFunc {
	checks := []health.Config{
		{
			Name:      "db",
			Timeout:   time.Second * 5,
			SkipOnErr: false,
			Check:     mongo.New(mongo.Config{DSN: conf.DB}),
		},
	}

	h, err := health.New(health.WithComponent(health.Component{
		Name:    "scaffold",
		Version: conf.Version,
	}), health.WithChecks(checks...))
	if err != nil {
		log.Fatalf("failed to create health check: %v", err)
	}

	return func(c echo.Context) error {
		// Optional HTTP Basic Auth
		if conf.HealthCheck.Username != "" && conf.HealthCheck.Password != "" {
			username, password, ok := c.Request().BasicAuth()
			if !ok || username != conf.HealthCheck.Username || password != conf.HealthCheck.Password {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "unauthorized",
				})
			}
		}

		// Serve the health check
		h.Handler().ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
