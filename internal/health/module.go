package health

import "github.com/alexfalkowski/go-signal/v2"

// Module registers health checks and HTTP observers for health, liveness, and readiness.
var Module = signal.Module(
	signal.Register(register),
	signal.Register(httpHealthObserver),
	signal.Register(httpLivenessObserver),
	signal.Register(httpReadinessObserver),
)
