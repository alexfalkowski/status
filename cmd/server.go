package cmd

import (
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"github.com/alexfalkowski/status/config"
	"github.com/alexfalkowski/status/server/health"
	v1 "github.com/alexfalkowski/status/server/v1"
	"github.com/alexfalkowski/status/transport"
	"go.uber.org/fx"
)

// ServerOptions for cmd.
var ServerOptions = []fx.Option{
	fx.NopLogger, runtime.Module, debug.Module, feature.Module,
	telemetry.Module, metrics.Module, Module,
	config.Module, health.Module, transport.Module, v1.Module,
}
