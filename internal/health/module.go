package health

import "github.com/alexfalkowski/go-service/v2/di"

// Module registers health checks and HTTP observers for health, liveness, and readiness.
var Module = di.Module(
	di.Register(register),
	di.Register(httpHealthObserver),
	di.Register(httpLivenessObserver),
	di.Register(httpReadinessObserver),
)
