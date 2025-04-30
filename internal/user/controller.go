package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	ctxUtil "server/util/ctx"
	httpUtil "server/util/http"
)

type UserController struct {
	repository *UserRepository
}

func NewUserController(repository *UserRepository) *UserController {
	return &UserController{
		repository: repository,
	}
}

// Handles the GET request for a user
func (c *UserController) GetUser(w http.ResponseWriter, req *http.Request) {
	logger := ctxUtil.GetLogger(req.Context())

	idStr := chi.URLParam(req, "id")

	// Convert string to UUID
	id, err := uuid.Parse(idStr)
	if err != nil {
		httpUtil.BadRequest(w, "Invalid UUID format")
		return
	}

	userEntity, err := c.repository.Get(id)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			httpUtil.NotFound(w, "No user exists with this ID")
			return
		}

		logger.Log().Err(err).Msg("Failed fetching user")
		httpUtil.InternalServerError(w, "An error occured while fetching the user")
		return
	}

	// Return the user as JSON
	json.NewEncoder(w).Encode(userEntity.ToDTO())
}
