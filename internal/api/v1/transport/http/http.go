package http

import (
	"net/url"
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

// ErrInvalidLocation is returned when the requested redirect location is unsupported.
var ErrInvalidLocation = errors.New("status: invalid location")

// Response marks the status route response type for the shared REST transport.
type Response any

// Register adds GET /v1/status/{code} to the shared REST router.
//
// The route accepts status codes from 200 through 999. The optional sleep query
// parameter is parsed as a duration, rejected when it exceeds the configured
// maximum, and returns 408 Request Timeout when the request context is canceled
// while waiting. The optional location query parameter sets a Location response
// header for 3xx status codes.
func Register(cfg *config.Config) {
	rest.Get("/v1/status/{code}", func(ctx context.Context) (*Response, error) {
		req := meta.Request(ctx)
		res := meta.Response(ctx)

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

		if location := req.URL.Query().Get("location"); !strings.IsEmpty(location) {
			if !isRedirectStatusCode(code) || !isValidLocation(location) {
				return nil, status.SafeError(http.StatusBadRequest, ErrInvalidLocation)
			}

			res.Header().Set("Location", location)
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

func isRedirectStatusCode(code int) bool {
	return code >= http.StatusMultipleChoices && code < http.StatusBadRequest
}

func isValidLocation(location string) bool {
	if strings.Contains(location, "\r") || strings.Contains(location, "\n") {
		return false
	}

	_, err := url.Parse(location)
	return err == nil
}
