package cmd

import (
	"github.com/alexfalkowski/go-service/v2/cli"
	"github.com/alexfalkowski/go-service/v2/debug"
	"github.com/alexfalkowski/go-service/v2/feature"
	"github.com/alexfalkowski/go-service/v2/module"
	"github.com/alexfalkowski/go-service/v2/telemetry"
	"github.com/alexfalkowski/go-service/v2/transport"
	v1 "github.com/alexfalkowski/status/internal/api/v1"
	"github.com/alexfalkowski/status/internal/config"
	"github.com/alexfalkowski/status/internal/health"
)

// RegisterServer for cmd.
func RegisterServer(command cli.Commander) {
	cmd := command.AddServer("server", "Start status server",
		module.Module, debug.Module, feature.Module,
		telemetry.Module, transport.Module,
		health.Module, config.Module,
		v1.Module, cli.Module,
	)
	cmd.AddInput("")
}
