package http

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"

	shttp "github.com/alexfalkowski/go-service/transport/http"
)

// Register for http.
func Register(server *shttp.Server) error {
	return server.Mux.HandlePath("GET", "/v1/status/{code}", func(w http.ResponseWriter, r *http.Request, p map[string]string) {
		s := r.URL.Query().Get("sleep")

		if s != "" {
			t, _ := time.ParseDuration(s)
			time.Sleep(t)
		}

		c, err := strconv.Atoi(p["code"])
		if err != nil {
			c = http.StatusInternalServerError
		}

		req, err := httputil.DumpRequest(r, true)
		if err != nil {
			c = http.StatusInternalServerError
		}

		w.WriteHeader(c)
		w.Write(req)                                                 //nolint:errcheck
		w.Write([]byte(fmt.Sprintf("%d %s", c, http.StatusText(c)))) //nolint:errcheck
	})
}
