package api

import (
	"net/http"
	"strings"

	"server/internal/api"
	"server/internal/domain/auth"
	"server/internal/domain/auth/user/service"
)

const (
	HeaderKeyAuthorization = "Authorization"
)

// Middleware that authenticates the user and sets the user in the request context
func AuthenticatedMiddleware(authService *service.UserAuthService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorizationHeader := r.Header.Get(HeaderKeyAuthorization)

			// If authorization header is not provided, return unauthorized
			if authorizationHeader == "" {
				api.Unauthorized(w, "Authorization header is required")
				return
			}

			// Parse bearer token from authorization header
			token := strings.TrimPrefix(authorizationHeader, "Bearer ")
			if token == authorizationHeader {
				api.Unauthorized(w, "Invalid authorization header")
				return
			}

			// Verify token
			userEntity, tokenEntity, err := authService.VerifyAccessToken(token)
			if err != nil {
				api.Unauthorized(w, "Invalid token")
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
