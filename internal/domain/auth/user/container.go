package user

import (
	"server/internal/api/router"
	"server/internal/domain/auth/user/api"
	"server/internal/domain/auth/user/repository"
	"server/internal/domain/auth/user/service"

	"github.com/go-chi/chi/v5"
)

type UserAuthDepsContainer struct {
	Repository *repository.TokenRepository
	Service    *service.UserAuthService
	Controller *api.UserAuthController
}

func NewUserAuthDepsContainer(
	repository *repository.TokenRepository,
	service *service.UserAuthService,
	controller *api.UserAuthController,
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
				r.Use(api.AuthenticatedMiddleware(c.Service))

				r.Get("/info", c.Controller.Info)
				r.Post("/logout", c.Controller.Logout)
			})
		})
	}
}
