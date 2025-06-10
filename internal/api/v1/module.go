package v1

import (
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/status/internal/api/v1/transport/http"
)

// Module for fx.
var Module = di.Module(
	di.Register(http.Register),
)
