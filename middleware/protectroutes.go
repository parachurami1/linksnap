package middleware

import (
	"context"
	"linksnap/auth"
	"net/http"
	"strings"
)

func Protect(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		tokn, err := auth.ParseTokenClaims(bearer)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "claims", tokn)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
