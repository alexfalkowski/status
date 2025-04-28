package http

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/alexfalkowski/go-service/net/http/meta"
	"github.com/alexfalkowski/go-service/net/http/rest"
	"github.com/alexfalkowski/go-service/net/http/status"
	"github.com/alexfalkowski/go-service/strings"
)

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

		c, err := strconv.Atoi(req.PathValue("code"))
		if err != nil {
			return nil, status.Error(http.StatusBadRequest, err.Error())
		}

		return nil, status.Error(c, fmt.Sprintf("%d %s", c, http.StatusText(c)))
	})
}
