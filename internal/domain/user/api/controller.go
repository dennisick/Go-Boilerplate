package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"server/internal/api"
	"server/internal/core"
	"server/internal/domain/user/repository"
)

type UserController struct {
	repository *repository.UserRepository
}

func NewUserController(repository *repository.UserRepository) *UserController {
	return &UserController{
		repository: repository,
	}
}

// Handles the GET request for a user
func (c *UserController) GetUser(w http.ResponseWriter, req *http.Request) {
	logger := core.GetLogger(req.Context())

	idStr := chi.URLParam(req, "id")

	// Convert string to UUID
	id, err := uuid.Parse(idStr)
	if err != nil {
		api.BadRequest(w, "Invalid UUID format")
		return
	}

	userEntity, err := c.repository.Get(id)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			api.NotFound(w, "No user exists with this ID")
			return
		}

		logger.Log().Err(err).Msg("Failed fetching user")
		api.InternalServerError(w, "An error occured while fetching the user")
		return
	}

	// Return the user as JSON
	json.NewEncoder(w).Encode(userEntity.ToDTO())
}
