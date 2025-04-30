package application

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ApplicationRepository struct {
	db *pgxpool.Pool
}

func NewApplicationRepository(db *pgxpool.Pool) *ApplicationRepository {
	return &ApplicationRepository{
		db: db,
	}
}

// Gets an application by its ID from the database
func (r *ApplicationRepository) Get(id uuid.UUID) (*Application, error) {
	query := `SELECT id, owner_id, name, connectors, created_at FROM applications WHERE id = $1`

	var application Application

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := r.db.QueryRow(ctx, query, id).Scan(
		&application.ID,
		&application.OwnerID,
		&application.Name,
		&application.Connectors,
		&application.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrApplicationNotFound
		}

		return nil, fmt.Errorf("database error: %w", err)
	}

	return &application, nil
}
