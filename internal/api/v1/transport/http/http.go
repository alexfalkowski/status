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
)

// ErrInvalidStatusCode is returned when the requested status code is unsupported.
var ErrInvalidStatusCode = errors.New("status: invalid status code")

// Response for route.
type Response any

// Register for http.
func Register() {
	rest.Get("/v1/status/{code}", func(ctx context.Context) (*Response, error) {
		req := meta.Request(ctx)
		query := req.URL.Query()

		if s := query.Get("sleep"); !strings.IsEmpty(s) {
			t, err := time.ParseDuration(s)
			if err != nil {
				return nil, status.SafeError(http.StatusBadRequest, err)
			}

			timer := time.NewTimer(t)
			defer timer.Stop()

			select {
			case <-ctx.Done():
				return nil, status.SafeError(http.StatusRequestTimeout, ctx.Err())
			case <-timer.C:
			}
		}

		code, err := parseStatusCode(req.PathValue("code"))
		if err != nil {
			return nil, status.SafeError(http.StatusBadRequest, err)
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
