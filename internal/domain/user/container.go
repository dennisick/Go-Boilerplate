package user

import (
	"server/internal/domain/user/api"
	"server/internal/domain/user/repository"

	"github.com/go-chi/chi/v5"
)

type UserDepsContainer struct {
	Repository *repository.UserRepository
	Controller *api.UserController
}

func NewUserDepsContainer(
	userRepository *repository.UserRepository,
	userController *api.UserController,
) *UserDepsContainer {
	return &UserDepsContainer{
		Repository: userRepository,
		Controller: userController,
	}
}

// Registers the routes for the user domain
func (c *UserDepsContainer) RegisterRoutes(r chi.Router) func(r chi.Router) {
	return func(r chi.Router) {
		r.Route("/v1/users", func(r chi.Router) {
			r.Get("/{id}", c.Controller.GetUser)
		})
	}
}
