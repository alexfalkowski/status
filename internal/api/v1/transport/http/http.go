package http

import (
	"math"
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

// ErrInvalidRetryAfter is returned when the requested retry after duration is unsupported.
var ErrInvalidRetryAfter = errors.New("status: invalid retry after")

// Response marks the status route response type for the shared REST transport.
type Response any

// Register adds /v1/status/{code} routes to the shared REST router.
//
// The routes accept status codes from 200 through 999. The optional sleep
// query parameter is parsed as a duration, rejected when it exceeds the
// configured maximum, and returns 408 Request Timeout when the request context
// is canceled while waiting. The optional location query parameter sets a
// Location response header for 3xx status codes. The optional retry_after query
// parameter sets a Retry-After response header for 3xx, 429, and 503 status
// codes.
func Register(cfg *config.Config) {
	for _, method := range []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete} {
		rest.Route(strings.Join(strings.Space, method, "/v1/status/{code}"), statusHandler(cfg))
	}
}

func statusHandler(cfg *config.Config) func(context.Context) (*Response, error) {
	return func(ctx context.Context) (*Response, error) {
		req := meta.Request(ctx)
		res := meta.Response(ctx)

		code, err := parseStatusCode(req.PathValue("code"))
		if err != nil {
			return nil, status.SafeError(http.StatusBadRequest, err)
		}

		if err := waitForSleep(ctx, cfg, req.URL.Query().Get("sleep")); err != nil {
			return nil, err
		}

		if err := setLocationHeader(res, code, req.URL.Query().Get("location")); err != nil {
			return nil, err
		}

		if err := setRetryAfterHeader(res, code, req.URL.Query().Get("retry_after")); err != nil {
			return nil, err
		}

		return nil, status.Errorf(code, "%d %s", code, http.StatusText(code))
	}
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

func waitForSleep(ctx context.Context, cfg *config.Config, sleepValue string) error {
	if strings.IsEmpty(sleepValue) {
		return nil
	}

	sleep, err := time.ParseDuration(sleepValue)
	if err != nil {
		return status.SafeError(http.StatusBadRequest, err)
	}
	if sleep > cfg.GetMaxSleep() {
		return status.SafeError(http.StatusBadRequest, ErrInvalidSleepDuration)
	}

	timer := time.NewTimer(sleep)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return status.SafeError(http.StatusRequestTimeout, ctx.Err())
	case <-timer.C:
		return nil
	}
}

func setLocationHeader(res http.ResponseWriter, code int, location string) error {
	if strings.IsEmpty(location) {
		return nil
	}

	if !isRedirectStatusCode(code) || !isValidLocation(location) {
		return status.SafeError(http.StatusBadRequest, ErrInvalidLocation)
	}

	res.Header().Set("Location", location)
	return nil
}

func setRetryAfterHeader(res http.ResponseWriter, code int, retryAfterValue string) error {
	if strings.IsEmpty(retryAfterValue) {
		return nil
	}

	if !isRetryAfterStatusCode(code) {
		return status.SafeError(http.StatusBadRequest, ErrInvalidRetryAfter)
	}

	retryAfter, err := time.ParseDuration(retryAfterValue)
	if err != nil {
		return status.SafeError(http.StatusBadRequest, ErrInvalidRetryAfter)
	}
	if retryAfter <= 0 {
		return status.SafeError(http.StatusBadRequest, ErrInvalidRetryAfter)
	}

	seconds := int64(math.Ceil(retryAfter.Duration().Seconds()))
	res.Header().Set("Retry-After", strconv.FormatInt(seconds, 10))
	return nil
}

func isRedirectStatusCode(code int) bool {
	return code >= http.StatusMultipleChoices && code < http.StatusBadRequest
}

func isRetryAfterStatusCode(code int) bool {
	return isRedirectStatusCode(code) ||
		code == http.StatusTooManyRequests ||
		code == http.StatusServiceUnavailable
}

func isValidLocation(location string) bool {
	if strings.Contains(location, "\r") || strings.Contains(location, "\n") {
		return false
	}

	_, err := url.Parse(location)
	return err == nil
}
