package middleware

import (
	"log/slog"
	"net/http"
)

func LogRequest(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Incoming Request", "host", r.RemoteAddr, "enpoint", r.URL.Path)
		h.ServeHTTP(w, r)
	}
}
