package http

import (
	"fmt"
	"net/http"
	"strconv"

	shttp "github.com/alexfalkowski/go-service/transport/http"
)

// Register for http.
func Register(server *shttp.Server) error {
	return server.Mux.HandlePath("GET", "/v1/status/{code}", func(w http.ResponseWriter, r *http.Request, p map[string]string) {
		s, err := strconv.Atoi(p["code"])
		if err != nil {
			s = http.StatusInternalServerError
		}

		w.WriteHeader(s)
		w.Write([]byte(fmt.Sprintf("%d %s", s, http.StatusText(s)))) //nolint:errcheck
	})
}
