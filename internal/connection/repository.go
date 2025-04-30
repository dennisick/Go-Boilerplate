package connection

import (
	"context"
	"errors"
	"fmt"
	"server/internal/application"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ConnectionRepository struct {
	db *pgxpool.Pool
}

func NewConnectionRepository(db *pgxpool.Pool) *ConnectionRepository {
	return &ConnectionRepository{
		db: db,
	}
}

// Returns a connection with the application it belongs to
// If the connection does not exist, the error ErrConnectionNotFound is returned
func (r *ConnectionRepository) GetWithApplication(id uuid.UUID) (*Connection, error) {
	query := `
		SELECT c.id, a.id, c.email, c.provider, c.refresh_token, c.created_at, 
		a.id, a.owner_id, a.name, a.connectors, a.created_at
		FROM connections c
		INNER JOIN applications a ON a.id = c.application_id
		WHERE c.id = $1
	`

	var connection Connection
	var application application.Application

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := r.db.QueryRow(ctx, query, id).Scan(
		&connection.ID,
		&connection.ApplicationID,
		&connection.Email,
		&connection.Provider,
		&connection.RefreshToken,
		&connection.CreatedAt,
		&application.ID,
		&application.OwnerID,
		&application.Name,
		&application.Connectors,
		&application.CreatedAt,
	)

	if err != nil {
		// Check if it's a "no rows" error
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrConnectionNotFound
		}

		// For other database errors, return the original error
		return nil, fmt.Errorf("database error: %w", err)
	}

	connection.Application = &application

	return &connection, nil
}

func (r *ConnectionRepository) Delete(id uuid.UUID) (bool, error) {
	statement := `DELETE FROM connections WHERE id= $1`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	res, err := r.db.Exec(ctx, statement, id)
	if err != nil {
		return false, fmt.Errorf("database error: %w", err)
	}

	return res.RowsAffected() > 0, nil
}
