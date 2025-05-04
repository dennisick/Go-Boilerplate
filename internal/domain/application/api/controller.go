package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"server/internal/api"
	"server/internal/core"
	"server/internal/domain/application/repository"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ApplicationController struct {
	repository *repository.ApplicationRepository
}

func NewApplicationController(repository *repository.ApplicationRepository) *ApplicationController {
	return &ApplicationController{
		repository: repository,
	}
}

// Handles the GET request for the application by ID
func (c *ApplicationController) GetApplication(w http.ResponseWriter, req *http.Request) {
	logger := core.GetLogger(req.Context())

	idStr := chi.URLParam(req, "id")

	// Convert string to UUID
	id, err := uuid.Parse(idStr)
	if err != nil {
		api.BadRequest(w, "Invalid UUID format")
		return
	}

	// Get the application from the repository
	application, err := c.repository.Get(id)
	if err != nil {
		// If the application is not found, return a 404
		if errors.Is(err, repository.ErrApplicationNotFound) {
			api.NotFound(w, "No application exists with this ID")
			return
		}

		// Otherwise, log the error and return a 500
		logger.Log().Err(err).Msg("Failed fetching application")
		api.InternalServerError(w, "An error occured while fetching the application")
		return
	}

	// Return the application as JSON
	json.NewEncoder(w).Encode(application.ToDTO())
}
