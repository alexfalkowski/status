package cmd

import (
	"github.com/alexfalkowski/status/config"
	"github.com/alexfalkowski/status/server/health"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"github.com/alexfalkowski/go-service/transport"
	"go.uber.org/fx"
)

// ServerOptions for cmd.
var ServerOptions = []fx.Option{
	fx.NopLogger, runtime.Module,
	telemetry.Module, metrics.Module, Module,
	config.Module, health.Module, transport.Module,
}
