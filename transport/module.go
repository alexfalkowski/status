package transport

import (
	"github.com/alexfalkowski/go-service/transport"
	"github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/http"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	transport.Module,
	fx.Provide(http.ServerHandlers),
	fx.Provide(grpc.UnaryServerInterceptor),
	fx.Provide(grpc.StreamServerInterceptor),
)
