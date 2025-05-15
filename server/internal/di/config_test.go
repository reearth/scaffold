package di

import (
	"testing"

	"github.com/reearth/reearthx/appx"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	t.Setenv("PORT", "8081")
	t.Setenv("REEARTH_DB", "test")
	t.Setenv("REEARTH_AUTH_ISS", "https://example.com")
	t.Setenv("REEARTH_AUTH_JWKSURI", "https://example.com/jwks.json")
	t.Setenv("REEARTH_AUTH_AUD", "aud1,aud2")
	t.Setenv("REEARTH_AUTH_TTL", "10")
	t.Setenv("REEARTH_DEV", "true")

	cfg := loadConfig(false)
	assert.Equal(t, &Config{
		Dev:    true,
		DB:     "test",
		DB_APP: "reearth",
		Port:   "8081",
		Auth: appx.JWTProvider{
			ISS:     "https://example.com",
			JWKSURI: lo.ToPtr("https://example.com/jwks.json"),
			AUD:     []string{"aud1", "aud2"},
			TTL:     lo.ToPtr(10),
		},
	}, cfg)
}
