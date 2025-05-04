package middleware

import (
	"net/http"
	"server/internal/api"
	"time"
)

// Middleware that adds the request time to the request's context
func RequestTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = api.SetRequestTime(ctx, time.Now())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
