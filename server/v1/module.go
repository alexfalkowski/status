package v1

import (
	"github.com/alexfalkowski/status/server/v1/transport/http"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Invoke(http.Register),
)
