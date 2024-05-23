package http

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	sh "github.com/alexfalkowski/go-service/net/http"
)

// Register for http.
func Register(mux sh.ServeMux) error {
	return mux.Handle("GET", "/v1/status/{code}", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		s := query.Get("sleep")

		if s != "" {
			t, err := time.ParseDuration(s)
			if err != nil {
				writeBadRequest(w, err)

				return
			}

			time.Sleep(t)
		}

		c, err := strconv.Atoi(r.PathValue("code"))
		if err != nil {
			writeBadRequest(w, err)

			return
		}

		w.WriteHeader(c)
		w.Write([]byte(fmt.Sprintf("%d %s", c, http.StatusText(c)))) //nolint:errcheck
	})
}

func writeBadRequest(w http.ResponseWriter, err error) {
	writeError(w, http.StatusBadRequest, err)
}

func writeError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	w.Write([]byte(err.Error())) //nolint:errcheck
}
