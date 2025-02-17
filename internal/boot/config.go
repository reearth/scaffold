package boot

import (
	"os"

	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/yudai/pp"
)

type Config struct {
	DB     string `env:"DB" pp:"-"`
	DB_APP string `env:"DB_APP" default:"reearth"`
	Port   string `env:"PORT" envDefault:"8080"`
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

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
