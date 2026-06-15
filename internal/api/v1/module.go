package v1

import (
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/status/internal/api/v1/transport/http"
)

// Module registers the v1 HTTP API routes.
var Module = di.Module(
	di.Register(http.Register),
)
