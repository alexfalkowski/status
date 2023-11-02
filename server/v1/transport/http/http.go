package http

import (
	"net/http"
	"strconv"

	shttp "github.com/alexfalkowski/go-service/transport/http"
)

// Register for http.
func Register(server *shttp.Server) error {
	return server.Mux.HandlePath("GET", "/v1/status/{code}", func(w http.ResponseWriter, r *http.Request, p map[string]string) {
		c, err := strconv.Atoi(p["code"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(c)
		}
	})
}
