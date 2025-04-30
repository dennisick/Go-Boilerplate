package user

import (
	"errors"
	"fmt"
	"server/internal/auth"
	"server/internal/config"
	"server/internal/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserAuthService struct {
	config          *config.GeneralConfig
	userRepository  *user.UserRepository
	tokenRepository *TokenRepository
}

func NewUserAuthService(config *config.ApplicationConfig, up *user.UserRepository, tp *TokenRepository) *UserAuthService {
	return &UserAuthService{
		config:          config.General,
		userRepository:  up,
		tokenRepository: tp,
	}
}

// Verifies a user by their email and password and returns the user entity
func (s *UserAuthService) VerifyUser(email string, password string) (*user.User, error) {
	userEntity, err := s.userRepository.GetByEmail(email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrUserNotFound
		}

		return nil, fmt.Errorf("database error: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userEntity.Password), []byte(password))
	if err != nil {
		return nil, auth.ErrInvalidCredentials
	}

	return userEntity, nil
}

// Verifies an access token and returns the user entity and the token entity
func (s *UserAuthService) VerifyAccessToken(encryptedToken string) (*user.User, *auth.UserToken, error) {
	// Parse and validate token
	token, err := jwt.Parse(encryptedToken, func(token *jwt.Token) (interface{}, error) {
		return s.config.SecretKey, nil
	})

	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, nil, auth.ErrInvalidToken
	}

	// Parse token ID and user ID to UUID
	var tokenID uuid.UUID
	var userID uuid.UUID

	if tokenID, err = uuid.Parse(claims["jti"].(string)); err != nil {
		return nil, nil, auth.ErrInvalidToken
	}
	if userID, err = uuid.Parse(claims["sub"].(string)); err != nil {
		return nil, nil, auth.ErrInvalidToken
	}

	// Get token entity from database
	tokenEntity, err := s.tokenRepository.GetAccessTokenByID(tokenID)
	if err != nil {
		return nil, nil, err
	}

	// Check if token belongs to user
	if tokenEntity.SubjectID != userID {
		return nil, nil, auth.ErrInvalidToken
	}

	// Get user entity from database
	user, err := s.userRepository.Get(userID)
	if err != nil {
		return nil, nil, err
	}

	return user, tokenEntity, nil
}

// Verifies a refresh token and returns the access token entity and the refresh token entity
func (s *UserAuthService) VerifyRefreshToken(encryptedToken string) (*auth.UserToken, *auth.UserToken, error) {
	// Parse and validate token
	token, err := jwt.Parse(encryptedToken, func(token *jwt.Token) (interface{}, error) {
		return s.config.SecretKey, nil
	})

	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, nil, auth.ErrInvalidToken
	}

	// Parse token ID and user ID to UUID
	var refreshTokenID uuid.UUID
	var accessTokenID uuid.UUID

	if refreshTokenID, err = uuid.Parse(claims["jti"].(string)); err != nil {
		return nil, nil, auth.ErrInvalidToken
	}
	if accessTokenID, err = uuid.Parse(claims["sub"].(string)); err != nil {
		return nil, nil, auth.ErrInvalidToken
	}

	// Get token entity from database
	refreshTokenEntity, err := s.tokenRepository.GetRefreshTokenByID(refreshTokenID)
	if err != nil {
		return nil, nil, err
	}

	// Check if token belongs to user
	if refreshTokenEntity.SubjectID != accessTokenID {
		return nil, nil, auth.ErrInvalidToken
	}

	// Get access token entity from database
	accessTokenEntity, err := s.tokenRepository.GetAccessTokenByID(accessTokenID)
	if err != nil {
		return nil, nil, err
	}

	return accessTokenEntity, refreshTokenEntity, nil
}

// Refreshes the access token and returns the updated access token entity
func (s *UserAuthService) RefreshAccessToken(accessToken *auth.UserToken) (*auth.UserToken, error) {
	// Define new expiry time
	newExpiryTime := time.Now().Add(auth.AccessTokenExpiry)

	// Update access token in database
	_, err := s.tokenRepository.UpdateAccessToken(accessToken.ID, newExpiryTime)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	// Update access token expiry time
	accessToken.ExpiresAt = newExpiryTime

	return accessToken, nil
}
