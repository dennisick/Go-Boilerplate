package application

import (
	"server/internal/api/router"
	"server/internal/domain/application/api"
	"server/internal/domain/application/repository"

	"github.com/go-chi/chi/v5"
)

type ApplicationDepsContainer struct {
	Repository *repository.ApplicationRepository
	Controller *api.ApplicationController
}

func NewApplicationDepsContainer(
	applicationRepository *repository.ApplicationRepository,
	applicationController *api.ApplicationController,
) *ApplicationDepsContainer {
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
