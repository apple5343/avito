package config

import "github.com/ilyakaznacheev/cleanenv"

type Postgres struct {
	DSN string `env:"POSTGRES_DSN" env-default:"postgres://postgres:postgres@localhost:5432/avito?sslmode=disable"`
}

func PostgresConfig() (*Postgres, error) {
	var cfg Postgres
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
