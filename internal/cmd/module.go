package cmd

import (
	"github.com/alexfalkowski/go-service/v3/module"
	"github.com/alexfalkowski/go-signal/v2"
	v1 "github.com/alexfalkowski/status/v2/internal/api/v1"
	"github.com/alexfalkowski/status/v2/internal/config"
	"github.com/alexfalkowski/status/v2/internal/health"
)

// Module composes the server command with service config, health, and v1 API wiring.
var Module = signal.Module(
	module.Server,
	config.Module,
	health.Module,
	v1.Module,
)
