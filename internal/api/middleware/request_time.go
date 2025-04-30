package middleware

import (
	"net/http"
	"time"

	ctxUtil "server/util/ctx"
)

// Middleware that adds the request time to the request's context
func RequestTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = ctxUtil.SetRequestTime(ctx, time.Now())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
