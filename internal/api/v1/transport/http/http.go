package http

import (
	"strconv"

	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/errors"
	"github.com/alexfalkowski/go-service/v2/net/http"
	"github.com/alexfalkowski/go-service/v2/net/http/meta"
	"github.com/alexfalkowski/go-service/v2/net/http/rest"
	"github.com/alexfalkowski/go-service/v2/net/http/status"
	"github.com/alexfalkowski/go-service/v2/strings"
	"github.com/alexfalkowski/go-service/v2/time"
	"github.com/alexfalkowski/status/internal/config"
)

// ErrInvalidStatusCode is returned when the requested status code is unsupported.
var ErrInvalidStatusCode = errors.New("status: invalid status code")

// ErrInvalidSleepDuration is returned when the requested sleep duration is unsupported.
var ErrInvalidSleepDuration = errors.New("status: invalid sleep duration")

// Response marks the status route response type for the shared REST transport.
type Response any

// Register adds GET /v1/status/{code} to the shared REST router.
//
// The route accepts status codes from 200 through 999. The optional sleep query
// parameter is parsed as a duration, rejected when it exceeds the configured
// maximum, and returns 408 Request Timeout when the request context is canceled
// while waiting.
func Register(cfg *config.Config) {
	rest.Get("/v1/status/{code}", func(ctx context.Context) (*Response, error) {
		req := meta.Request(ctx)

		code, err := parseStatusCode(req.PathValue("code"))
		if err != nil {
			return nil, status.SafeError(http.StatusBadRequest, err)
		}

		if s := req.URL.Query().Get("sleep"); !strings.IsEmpty(s) {
			sleep, err := time.ParseDuration(s)
			if err != nil {
				return nil, status.SafeError(http.StatusBadRequest, err)
			}
			if sleep > cfg.GetMaxSleep() {
				return nil, status.SafeError(http.StatusBadRequest, ErrInvalidSleepDuration)
			}

			timer := time.NewTimer(sleep)
			defer timer.Stop()

			select {
			case <-ctx.Done():
				return nil, status.SafeError(http.StatusRequestTimeout, ctx.Err())
			case <-timer.C:
			}
		}

		return nil, status.Errorf(code, "%d %s", code, http.StatusText(code))
	})
}

func parseStatusCode(code string) (int, error) {
	codeValue, err := strconv.Atoi(code)
	if err != nil {
		return 0, err
	}

	if codeValue < 200 || codeValue > 999 {
		return 0, ErrInvalidStatusCode
	}

	return codeValue, nil
}
