package middleware

import (
	"context"
	"linksnap/cache"
	"net/http"
)

func PassCache(cac *cache.Cache, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "cache", cac)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
