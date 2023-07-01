package config

import (
	"flag"

	"github.com/caarlos0/env"
)

type Config struct {
	Address  string `env:"ADDRESS"`
	Database string `env:"DATABASE"`
	Debug    bool   `env:"DEBUG"`
}

func New() (Config, error) {
	var cfg Config

	cfg.CheckFlags()

	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c *Config) CheckFlags() {
	flag.StringVar(&c.Address, "a", "127.0.0.1:8080", "port to connect")
	flag.StringVar(&c.Database, "d", "solver", "database dsn")
	flag.BoolVar(&c.Debug, "debug", true, "logger debug mode")

	flag.Parse()
}
