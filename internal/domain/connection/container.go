package connection

import (
	"server/internal/api/router"
	authApi "server/internal/domain/auth/user/api"
	"server/internal/domain/auth/user/service"
	"server/internal/domain/connection/api"
	"server/internal/domain/connection/repository"

	"github.com/go-chi/chi/v5"
)

type ConnectionDepsContainer struct {
	Repository  *repository.ConnectionRepository
	Controller  *api.ConnectionController
	AuthService *service.UserAuthService
}

func NewConnectionDepsContainer(
	connectionRepository *repository.ConnectionRepository,
	connectionController *api.ConnectionController,
	authService *service.UserAuthService,
) *ConnectionDepsContainer {
	return &ConnectionDepsContainer{
		Repository:  connectionRepository,
		Controller:  connectionController,
		AuthService: authService,
	}
}

// Registers the routes for the connection domain
func (c *ConnectionDepsContainer) RegisterRoutes(r chi.Router) router.Route {
	return func(r chi.Router) {
		r.Route("/v1/connections", func(r chi.Router) {
			r.Use(authApi.AuthenticatedMiddleware(c.AuthService))

			r.Get("/{id}", c.Controller.GetConnection)
			r.Delete("/{id}", c.Controller.DeleteConnection)
		})
	}
}
