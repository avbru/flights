package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port  string `env:"PORT" env-default:"8080"`
	PgUrl string `env:"PG_URL" env-default:"postgres://postgres:postgres@localhost:5432/flights?sslmode=disable"`
}

func New() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
