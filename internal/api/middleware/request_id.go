package middleware

import (
	"net/http"
	"server/internal/api"

	"github.com/google/uuid"
)

const (
	HeaderKeyRequestId = "X-Request-Id"
)

// Middleware that adds a request ID to the request's context
// If the request ID is not present, it will be generated
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		requestID := r.Header.Get(HeaderKeyRequestId)
		if requestID == "" {
			requestID = uuid.New().String()
		}

		w.Header().Set(HeaderKeyRequestId, requestID)

		ctx = api.SetRequestID(ctx, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
