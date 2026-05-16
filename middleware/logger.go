package middleware

import (
	"log/slog"
	"net/http"
	"strings"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("", "IP address", strings.Split(r.RemoteAddr, ":")[0], "Path", r.URL.Path, "Method", r.Method)
		next.ServeHTTP(w, r)
	})
}
