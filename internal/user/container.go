package user

import (
	"github.com/go-chi/chi/v5"
)

type UserDepsContainer struct {
	Repository *UserRepository
	Controller *UserController
}

func NewUserDepsContainer(userRepository *UserRepository, userController *UserController) *UserDepsContainer {
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
