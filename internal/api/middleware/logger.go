package middleware

import (
	"net/http"
	"server/internal/api"
	"server/internal/core"

	"github.com/rs/zerolog"
)

// Middleware that adds a logger to the request's context and
// appends the request's path and request ID to the logger
func Logger(l zerolog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			// Get request data and provide context in logger
			path := r.URL.EscapedPath()
			requestID := api.GetRequestID(ctx)

			logger := l.With().Timestamp().Str("path", path).Str("request_id", requestID).Logger()

			ctx = core.SetLogger(ctx, logger)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
