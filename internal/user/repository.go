package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Returns a user by their email
// If the user does not exist, the error ErrUserNotFound is returned
func (r *UserRepository) GetByEmail(email string) (*User, error) {
	query := `SELECT id, first_name, last_name, email, password, created_at FROM users WHERE email = $1`

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("database error: %w", err)
	}

	return &user, nil
}

// Returns a user by their ID
// If the user does not exist, the error ErrUserNotFound is returned
func (r *UserRepository) Get(id uuid.UUID) (*User, error) {
	query := `SELECT id, first_name, last_name, email, password, created_at FROM users WHERE id = $1`

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		// Check if it's a "no rows" error
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		// For other database errors, return the original error
		return nil, fmt.Errorf("database error: %w", err)
	}

	return &user, nil
}
