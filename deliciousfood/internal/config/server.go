package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Server struct {
	Addr     string `env:"DELICIOUSFOOD_ADDR" env-default:"0.0.0.0:8081"`
	AvitoURL string `env:"AVITO_URL" env-required:"true"`
	APIKey   string `env:"API_KEY" env-required:"true"`
}

func Load() (*Server, error) {
	cfg := &Server{}
	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("read env: %w", err)
	}
	return cfg, nil
}
