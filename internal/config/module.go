package config

import (
	"github.com/alexfalkowski/go-service/v3/config"
	"github.com/alexfalkowski/go-signal/v2"
)

// Module registers service config loading, validation, and shared config projection.
var Module = signal.Module(
	signal.Constructor(newValidation),
	signal.Constructor(config.NewConfig[Config]),
	signal.Decorate(decorateConfig),
	signal.Constructor(healthConfig),
)
