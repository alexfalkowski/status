package config

import (
	"github.com/alexfalkowski/status/health"
	"github.com/alexfalkowski/go-service/config"
)

// Config for the service.
type Config struct {
	Health        health.Config `yaml:"health" json:"health" toml:"health"`
	config.Config `yaml:",inline" json:",inline" toml:",inline"`
}
