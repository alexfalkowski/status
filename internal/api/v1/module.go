package v1

import (
	"github.com/alexfalkowski/go-signal/v2"
	"github.com/alexfalkowski/status/internal/api/v1/transport/http"
)

// Module registers the v1 HTTP API routes.
var Module = signal.Module(
	signal.Register(http.Register),
)
