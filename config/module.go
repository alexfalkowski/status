package config

import (
	"github.com/alexfalkowski/go-service/config"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(NewConfigurator),
	config.ConfigModule,
	fx.Provide(healthConfig),
)
