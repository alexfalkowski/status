package cmd

import (
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/go-service/v2/module"
	v1 "github.com/alexfalkowski/status/internal/api/v1"
	"github.com/alexfalkowski/status/internal/config"
	"github.com/alexfalkowski/status/internal/health"
)

// Module composes the server command with service config, health, and v1 API wiring.
var Module = di.Module(
	module.Server,
	config.Module,
	health.Module,
	v1.Module,
)
