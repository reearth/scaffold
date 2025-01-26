package boot

import (
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/yudai/pp"
)

type Config struct {
	DB     string `pp:"-"`
	DB_APP string
	Port   string `env:"PORT" envDefault:"8080"`
}

func (cfg *Config) Print() {
	pp.Printf("config: %+v\n", cfg)
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
