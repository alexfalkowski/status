package health

import (
	health "github.com/alexfalkowski/go-health/server"
	http "github.com/alexfalkowski/go-service/v2/transport/http/health"
)

func httpHealthObserver(server *health.Server) *http.HealthObserver {
	return &http.HealthObserver{Observer: server.Observe("online")}
}

func httpLivenessObserver(server *health.Server) *http.LivenessObserver {
	return &http.LivenessObserver{Observer: server.Observe("noop")}
}

func httpReadinessObserver(server *health.Server) *http.ReadinessObserver {
	return &http.ReadinessObserver{Observer: server.Observe("noop")}
}
