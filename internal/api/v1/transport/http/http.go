package http

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/net/http"
	"github.com/alexfalkowski/go-service/v2/net/http/meta"
	"github.com/alexfalkowski/go-service/v2/net/http/rest"
	"github.com/alexfalkowski/go-service/v2/net/http/status"
	"github.com/alexfalkowski/go-service/v2/strings"
)

var errInvalidStatusCode = errors.New("invalid status code")

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
				return nil, status.Error(http.StatusBadRequest, err.Error())
			}

			time.Sleep(t)
		}

		code, err := parseStatusCode(req.PathValue("code"))
		if err != nil {
			return nil, status.Error(http.StatusBadRequest, err.Error())
		}

		return nil, status.Error(code, fmt.Sprintf("%d %s", code, http.StatusText(code)))
	})
}

func parseStatusCode(code string) (int, error) {
	codeValue, err := strconv.Atoi(code)
	if err != nil {
		return 0, err
	}

	if codeValue < 100 || codeValue > 999 {
		return 0, errInvalidStatusCode
	}

	return codeValue, nil
}
