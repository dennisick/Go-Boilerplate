package user

import (
	"net/http"
	"strings"

	"server/internal/auth"
	httpUtil "server/util/http"
)

const (
	HeaderKeyAuthorization = "Authorization"
)

// Middleware that authenticates the user and sets the user in the request context
func AuthenticatedMiddleware(authService *UserAuthService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorizationHeader := r.Header.Get(HeaderKeyAuthorization)

			// If authorization header is not provided, return unauthorized
			if authorizationHeader == "" {
				httpUtil.Unauthorized(w, "Authorization header is required")
				return
			}

			// Parse bearer token from authorization header
			token := strings.TrimPrefix(authorizationHeader, "Bearer ")
			if token == authorizationHeader {
				httpUtil.Unauthorized(w, "Invalid authorization header")
				return
			}

			// Verify token
			userEntity, tokenEntity, err := authService.VerifyAccessToken(token)
			if err != nil {
				httpUtil.Unauthorized(w, "Invalid token")
				return
			}

			// Set user in request context
			ctx := r.Context()
			ctx = auth.SetCtxUser(ctx, userEntity)
			ctx = auth.SetCtxAccessToken(ctx, tokenEntity)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
