package user

import (
	"context"
	"errors"
	"fmt"
	"server/internal/auth"
	"server/internal/user"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TokenRepository struct {
	db *pgxpool.Pool
}

func NewTokenRepository(db *pgxpool.Pool) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}

// Creates a new access token for an user in the database
func (r *TokenRepository) CreateAccessToken(user *user.User) (*auth.UserToken, error) {
	statement := `INSERT INTO access_tokens (user_id, expires_at) VALUES ($1, $2) RETURNING id, created_at`

	// Initialize the token entity with user ID and expiry time
	var tokenEntity auth.UserToken
	tokenEntity.SubjectID = user.ID
	tokenEntity.ExpiresAt = time.Now().Add(auth.AccessTokenExpiry)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := r.db.QueryRow(ctx, statement, user.ID, time.Now()).Scan(
		&tokenEntity.ID,
		&tokenEntity.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	return &tokenEntity, nil
}

// Creates a new refresh token for an access token in the database
func (r *TokenRepository) CreateRefreshToken(accessToken *auth.UserToken) (*auth.UserToken, error) {
	statement := `INSERT INTO refresh_tokens (access_token_id, expires_at) VALUES ($1, $2) RETURNING id, created_at`

	// Initialize the token entity with user ID and expiry time
	var tokenEntity auth.UserToken
	tokenEntity.SubjectID = accessToken.ID
	tokenEntity.ExpiresAt = time.Now().Add(auth.RefreshTokenExpiry)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := r.db.QueryRow(ctx, statement, accessToken.ID, time.Now()).Scan(
		&tokenEntity.ID,
		&tokenEntity.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	return &tokenEntity, nil
}

// Gets an access token from the database by its ID
func (r *TokenRepository) GetAccessTokenByID(id uuid.UUID) (*auth.UserToken, error) {
	query := `SELECT id, user_id, expires_at, created_at FROM access_tokens WHERE id = $1`

	var tokenEntity auth.UserToken

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := r.db.QueryRow(ctx, query, id).Scan(
		&tokenEntity.ID,
		&tokenEntity.SubjectID,
		&tokenEntity.ExpiresAt,
		&tokenEntity.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, auth.ErrTokenNotFound
		}

		return nil, fmt.Errorf("database error: %w", err)
	}

	return &tokenEntity, nil
}

// Gets a refresh token from the database by its ID
func (r *TokenRepository) GetRefreshTokenByID(id uuid.UUID) (*auth.UserToken, error) {
	query := `SELECT id, access_token_id, expires_at, created_at FROM refresh_tokens WHERE id = $1`

	var tokenEntity auth.UserToken

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := r.db.QueryRow(ctx, query, id).Scan(
		&tokenEntity.ID,
		&tokenEntity.SubjectID,
		&tokenEntity.ExpiresAt,
		&tokenEntity.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, auth.ErrTokenNotFound
		}

		return nil, fmt.Errorf("database error: %w", err)
	}

	return &tokenEntity, nil
}

// Deletes an access token from the database by its ID
func (r *TokenRepository) DeleteAccessToken(id uuid.UUID) (bool, error) {
	statement := `DELETE FROM access_tokens WHERE id= $1`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	res, err := r.db.Exec(ctx, statement, id)
	if err != nil {
		return false, fmt.Errorf("database error: %w", err)
	}

	return res.RowsAffected() > 0, nil
}

// Updates the expiry time of an access token in the database by its ID
func (r *TokenRepository) UpdateAccessToken(id uuid.UUID, expiresAt time.Time) (bool, error) {
	statement := `UPDATE access_tokens SET expires_at = $1 WHERE id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	res, err := r.db.Exec(ctx, statement, expiresAt, id)
	if err != nil {
		return false, fmt.Errorf("database error: %w", err)
	}

	return res.RowsAffected() > 0, nil
}

// Deletes a refresh token from the database by its ID
func (r *TokenRepository) DeleteRefreshToken(id uuid.UUID) (bool, error) {
	statement := `DELETE FROM refresh_tokens WHERE id= $1`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	res, err := r.db.Exec(ctx, statement, id)
	if err != nil {
		return false, fmt.Errorf("database error: %w", err)
	}

	return res.RowsAffected() > 0, nil
}
