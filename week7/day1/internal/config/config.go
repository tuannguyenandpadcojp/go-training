package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env      string `envconfig:"ENV" default:"local"`
	GRPCAddr string `envconfig:"GRPC_ADDR" default:"localhost:9090"`
	Debug    bool   `envconfig:"DEBUG" default:"false"`
}

func LoadConfig(path string) (*Config, error) {
	if err := godotenv.Load(path); err != nil {
		return nil, fmt.Errorf("failed to load %s file: %w", path, err)
	}
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	return cfg, nil
}
