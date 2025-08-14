package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Addr     string         `yaml:"addr" env:"addr" env-default:":8080"`
	NATS     NATSConfig     `yaml:"nats" env:"nats"`
	Postgres PostgresConfig `yaml:"postgres" env:"postgres"`
}

type NATSConfig struct {
	Url string `yaml:"url" env:"NATS_URL" env-required:"true"`
}

type PostgresConfig struct {
	DSN string `yaml:"dsn" env:"POSTGRES_URL" env-required:"true"`
}

func MustLoad() *Config {
	filepath := "config/config.yaml"

	var cfg Config
	err := cleanenv.ReadConfig(filepath, &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return &cfg
}
