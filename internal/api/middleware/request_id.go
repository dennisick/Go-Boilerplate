package middleware

import (
	"net/http"

	"github.com/google/uuid"

	ctxUtil "server/util/ctx"
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

		ctx = ctxUtil.SetRequestID(ctx, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
