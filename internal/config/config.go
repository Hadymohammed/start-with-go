package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
    DatabaseURL string `envconfig:"DATABASE_URL" default:"postgres://postgres:postgres@localhost:5432/notes?sslmode=disable"`
    APIKey      string `envconfig:"API_KEY" required:"true"`
    Port        string `envconfig:"PORT" default:"8080"`
}

func Load() (Config, error) {
    var cfg Config
    if err := envconfig.Process("", &cfg); err != nil {
        return Config{}, err
    }
    return cfg, nil
}