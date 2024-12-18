package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/status/health"
)

// Config for the service.
type Config struct {
	Health         *health.Config `yaml:"health,omitempty" json:"health,omitempty" toml:"health,omitempty"`
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
}

func decorateConfig(cfg *Config) *config.Config {
	return cfg.Config
}

func healthConfig(cfg *Config) *health.Config {
	return cfg.Health
}
