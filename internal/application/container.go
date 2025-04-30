package application

import (
	"server/internal/api/router"

	"github.com/go-chi/chi/v5"
)

type ApplicationDepsContainer struct {
	Repository *ApplicationRepository
	Controller *ApplicationController
}

func NewApplicationDepsContainer(applicationRepository *ApplicationRepository, applicationController *ApplicationController) *ApplicationDepsContainer {
	return &ApplicationDepsContainer{
		Repository: applicationRepository,
		Controller: applicationController,
	}
}

// Registers the routes for the application domain
func (c *ApplicationDepsContainer) RegisterRoutes(r chi.Router) router.Route {
	return func(r chi.Router) {
		r.Route("/v1/applications", func(r chi.Router) {
			r.Get("/{id}", c.Controller.GetApplication)
		})
	}
}
