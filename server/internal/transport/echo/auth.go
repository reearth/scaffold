package echo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/labstack/echo/v4"
	"github.com/reearth/scaffold/server/internal/transport/gql"
	"github.com/reearth/scaffold/server/internal/usecase"
)

type AuthConfig struct {
	ISS string
	AUD []string
}

func (a AuthConfig) Validator() (*validator.Validator, error) {
	issurl, err := url.Parse(a.ISS)
	if err != nil {
		return nil, fmt.Errorf("failed to parse iss: %w", err)
	}

	const cacheTTL = 5 * time.Minute

	p := jwks.NewCachingProvider(issurl, cacheTTL)
	return validator.New(
		p.KeyFunc,
		validator.RS256,
		a.ISS,
		a.AUD,
	)
}

func UseAuthMiddleware(e *echo.Group, config AuthConfig, uc usecase.Usecases) error {
	validator, err := config.Validator()
	if err != nil {
		return err
	}

	e.Use(
		echo.WrapMiddleware(jwtmiddleware.New(validator.ValidateToken).CheckJWT),
		AuthMiddleware(config, uc),
	)

	return nil
}

func AuthMiddleware(config AuthConfig, uc usecase.Usecases) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			claims, ok := ctx.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
			if !ok {
				return next(c)
			}

			// payload, err := json.Marshal(claims)
			// if err != nil {
			// 	return echo.NewHTTPError(http.StatusInternalServerError, "failed to marshal claims")
			// }

			u, err := uc.User.FindBySub.Execute(ctx, claims.RegisteredClaims.Subject)
			if err != nil {
				return echo.NewHTTPError(500, "failed to find user")
			}
			if u == nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "user not found")
			}

			ctx = context.WithValue(ctx, gql.UserKey{}, u)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
