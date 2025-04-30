package middleware

import (
	"net/http"

	"github.com/rs/zerolog"

	ctxUtil "server/util/ctx"
)

// Middleware that adds a logger to the request's context and
// appends the request's path and request ID to the logger
func Logger(l zerolog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			// Get request data and provide context in logger
			path := r.URL.EscapedPath()
			requestID := ctxUtil.GetRequestID(ctx)

			logger := l.With().Timestamp().Str("path", path).Str("request_id", requestID).Logger()

			ctx = ctxUtil.SetLogger(ctx, logger)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
