package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/status/health"
)

// Config for the service.
type Config struct {
	Health        health.Config `yaml:"health" json:"health" toml:"health"`
	config.Config `yaml:",inline" json:",inline" toml:",inline"`
}
