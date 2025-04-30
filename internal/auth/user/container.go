package user

import (
	"server/internal/api/router"

	"github.com/go-chi/chi/v5"
)

type UserAuthDepsContainer struct {
	Repository *TokenRepository
	Service    *UserAuthService
	Controller *UserAuthController
}

func NewUserAuthDepsContainer(
	repository *TokenRepository,
	service *UserAuthService,
	controller *UserAuthController,
) *UserAuthDepsContainer {
	return &UserAuthDepsContainer{
		Repository: repository,
		Service:    service,
		Controller: controller,
	}
}

// Registers the routes for the user authentication domain
func (c *UserAuthDepsContainer) RegisterRoutes(r chi.Router) router.Route {
	return func(r chi.Router) {
		r.Route("/v1/auth", func(r chi.Router) {
			r.Post("/login", c.Controller.Login)
			r.Post("/refresh-token", c.Controller.RefreshToken)

			r.Group(func(r chi.Router) {
				// User has to be authenticated to access the routes below
				r.Use(AuthenticatedMiddleware(c.Service))

				r.Get("/info", c.Controller.Info)
				r.Post("/logout", c.Controller.Logout)
			})
		})
	}
}
