package connection

import (
	"server/internal/api/router"
	authUser "server/internal/auth/user"

	"github.com/go-chi/chi/v5"
)

type ConnectionDepsContainer struct {
	Repository  *ConnectionRepository
	Controller  *ConnectionController
	AuthService *authUser.UserAuthService
}

func NewConnectionDepsContainer(
	connectionRepository *ConnectionRepository,
	connectionController *ConnectionController,
	authService *authUser.UserAuthService,
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
			r.Use(authUser.AuthenticatedMiddleware(c.AuthService))

			r.Get("/{id}", c.Controller.GetConnection)
			r.Delete("/{id}", c.Controller.DeleteConnection)
		})
	}
}
