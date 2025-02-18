package di

import (
	"os"

	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/yudai/pp"
)

type Config struct {
	DB     string `pp:"-"`
	DB_APP string `envDefault:"reearth"`
	Port   string `env:"PORT" envDefault:"8080"`
	Dev    bool
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

	cfg.Print()
	return &cfg
}
