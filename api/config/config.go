package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Env        string `env:"BLOG_ENV" envDefault:"local"`
	Port       int    `env:"PORT" envDefault:"80"`
	DBHost     string `env:"BLOG_DB_HOST" envDefault:"127.0.0.1"`
	DBPort     int    `env:"BLOG_DB_PORT" envDefault:"33306"`
	DBUser     string `env:"BLOG_DB_USER" envDefault:"user"`
	DBPassword string `env:"BLOG_DB_PASSWORD" envDefault:"pass"`
	DBName     string `env:"BLOG_DB_NAME" envDefault:"blog"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
