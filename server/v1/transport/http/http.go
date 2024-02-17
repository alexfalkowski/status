package http

import (
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
			t, err := time.ParseDuration(s)
			if err != nil {
				writeBadRequest(w, err)

				return
			}

			time.Sleep(t)
		}

		c, err := strconv.Atoi(p["code"])
		if err != nil {
			writeBadRequest(w, err)

			return
		}

		req, err := httputil.DumpRequest(r, true)
		if err != nil {
			writeInternalError(w, err)

			return
		}

		w.WriteHeader(c)
		w.Write(req) //nolint:errcheck
	})
}

func writeInternalError(w http.ResponseWriter, err error) {
	writeError(w, http.StatusInternalServerError, err)
}

func writeBadRequest(w http.ResponseWriter, err error) {
	writeError(w, http.StatusBadRequest, err)
}

func writeError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	w.Write([]byte(err.Error())) //nolint:errcheck
}
