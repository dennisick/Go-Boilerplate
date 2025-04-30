package connection

import (
	"encoding/json"
	"errors"
	"net/http"

	"server/internal/auth"
	ctxUtil "server/util/ctx"
	httpUtil "server/util/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ConnectionController struct {
	repository *ConnectionRepository
}

func NewConnectionController(repository *ConnectionRepository) *ConnectionController {
	return &ConnectionController{
		repository: repository,
	}
}

// Handles the GET request for a connection
// If the requesting user is not the owner of the application, the request will be forbidden
func (c *ConnectionController) GetConnection(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	logger := ctxUtil.GetLogger(ctx)
	user := auth.GetCtxUser(ctx)
	idStr := chi.URLParam(req, "id")

	// Convert string to UUID
	id, err := uuid.Parse(idStr)
	if err != nil {
		httpUtil.BadRequest(w, "Invalid UUID format")
		return
	}

	connection, err := c.repository.GetWithApplication(id)
	if err != nil {
		if errors.Is(err, ErrConnectionNotFound) {
			httpUtil.NotFound(w, "No connection exists with this ID")
			return
		}

		logger.Log().Err(err).Msg("Failed fetching connection")
		httpUtil.InternalServerError(w, "An error occured while fetching the connection")
		return
	}

	if connection.Application.OwnerID != user.ID {
		httpUtil.Forbidden(w, "Forbidden Resource")
		return
	}

	// Return the connection as JSON
	json.NewEncoder(w).Encode(connection.ToDTO())
}

// Handles the DELETE request for a connection
// If the requesting user is not the owner of the application, the request will be forbidden
func (c *ConnectionController) DeleteConnection(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	logger := ctxUtil.GetLogger(ctx)
	user := auth.GetCtxUser(ctx)

	idStr := chi.URLParam(req, "id")

	// Convert string to UUID
	id, err := uuid.Parse(idStr)
	if err != nil {
		httpUtil.BadRequest(w, "Invalid UUID format")
		return
	}

	connection, err := c.repository.GetWithApplication(id)
	if err != nil {
		if errors.Is(err, ErrConnectionNotFound) {
			httpUtil.NotFound(w, "No connection exists with this ID")
			return
		}

		logger.Log().Err(err).Msg("Failed fetching connection")
		httpUtil.InternalServerError(w, "An error occured while fetching the connection")
		return
	}

	if connection.Application.OwnerID != user.ID {
		httpUtil.Forbidden(w, "Forbidden Resource")
		return
	}

	deleted, err := c.repository.Delete(id)
	if err != nil {
		logger.Log().Err(err).Msg("Failed deleting connection")
		httpUtil.InternalServerError(w, "An error occured while deleting the connection")

		return
	}

	// Return status if connection was deleted
	json.NewEncoder(w).Encode(deleted)
}
