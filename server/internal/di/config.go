package di

import (
	"os"

	"log"

	"github.com/joho/godotenv"
	"github.com/k0kubun/pp/v3"
	"github.com/kelseyhightower/envconfig"
	"github.com/reearth/reearthx/appx"
	"github.com/samber/lo"
)

func init() {
	pp.Default.SetColoringEnabled(false)
	pp.Default.SetOmitEmpty(true)
}

const configPrefix = "REEARTH"

type Config struct {
	Dev    bool
	DB     string `pp:"-"`
	DB_APP string `default:"reearth"`
	Port   string `envconfig:"PORT" default:"8080"`
	Auth   appx.JWTProvider
}

func (cfg *Config) JWTProviders() []appx.JWTProvider {
	if cfg.Auth.IsEmpty() {
		return nil
	}
	return []appx.JWTProvider{cfg.Auth}
}

func (cfg *Config) Print() {
	pp.Printf("config: %+v\n", cfg)
}

func LoadConfig() *Config {
	if err := godotenv.Load(".env"); err != nil && !os.IsNotExist(err) {
		panic(err)
	} else if err == nil {
		log.Println("config: .env loaded")
	}

	return loadConfig(true)
}

func loadConfig(print bool) *Config {
	var cfg Config
	lo.Must0(envconfig.Process(configPrefix, &cfg))

	if print {
		cfg.Print()
	}
	return &cfg
}
