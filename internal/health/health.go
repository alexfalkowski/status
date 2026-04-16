package health

import (
	"github.com/alexfalkowski/go-health/v2/checker"
	"github.com/alexfalkowski/go-health/v2/server"
	"github.com/alexfalkowski/go-service/v2/env"
	"github.com/alexfalkowski/go-service/v2/health"
)

func register(name env.Name, srv *server.Server, cfg *Config) {
	regs := health.Registrations{
		server.NewRegistration("noop", cfg.Duration.Duration(), checker.NewNoopChecker()),
		server.NewOnlineRegistration(cfg.Timeout.Duration(), cfg.Duration.Duration()),
	}

	srv.Register(name.String(), regs...)
}

func httpHealthObserver(name env.Name, server *server.Server) error {
	return server.Observe(name.String(), "healthz", "online")
}

func httpLivenessObserver(name env.Name, server *server.Server) error {
	return server.Observe(name.String(), "livez", "noop")
}

func httpReadinessObserver(name env.Name, server *server.Server) error {
	return server.Observe(name.String(), "readyz", "noop")
}
