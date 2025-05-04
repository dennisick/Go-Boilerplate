package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator"

	"server/internal/api"
	"server/internal/config"
	"server/internal/core"
	"server/internal/domain/auth"
	"server/internal/domain/auth/user/repository"
	"server/internal/domain/auth/user/service"
	userRepo "server/internal/domain/user/repository"
)

type UserAuthController struct {
	Config     *config.GeneralConfig
	Service    *service.UserAuthService
	Repository *repository.TokenRepository
}

func NewUserAuthController(
	config *config.ApplicationConfig,
	service *service.UserAuthService,
	repository *repository.TokenRepository,
) *UserAuthController {
	return &UserAuthController{
		Config:     config.General,
		Service:    service,
		Repository: repository,
	}
}

// Handles the POST request for the user login
func (c *UserAuthController) Login(w http.ResponseWriter, req *http.Request) {
	// Define the request body structure and validation rules
	type UserLoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8,max=32"`
		Remember bool   `json:"remember"`
	}

	logger := core.GetLogger(req.Context())

	// Decode request body to JSON
	var loginRequest UserLoginRequest
	if err := json.NewDecoder(req.Body).Decode(&loginRequest); err != nil {
		logger.Log().Err(err).Msg("Failed to decode login request")
		api.BadRequest(w, "Invalid request body")
		return
	}

	// Validate request body
	validate := validator.New()
	if err := validate.Struct(loginRequest); err != nil {
		api.BadRequest(w, err.Error())
		return
	}

	// Verify user
	userEntity, err := c.Service.VerifyUser(loginRequest.Email, loginRequest.Password)

	// Handle errors
	if err != nil {
		if errors.Is(err, userRepo.ErrUserNotFound) || errors.Is(err, auth.ErrInvalidCredentials) {
			api.Unauthorized(w, "Invalid credentials")
			return
		}

		logger.Log().Err(err).Msg("Failed to verify user")
		api.InternalServerError(w, "Failed to verify user")
		return
	}

	// Create access token
	tokenEntity, err := c.Repository.CreateAccessToken(userEntity)
	if err != nil {
		logger.Log().Err(err).Msg("Failed to create access token")
		api.InternalServerError(w, "Failed to create access token")
		return
	}

	// Encrypt access token
	encryptedToken, err := tokenEntity.ToJWT(c.Config.SecretKey)
	if err != nil {
		logger.Log().Err(err).Msg("Failed to encrypt access token")
		api.InternalServerError(w, "Failed to encrypt access token")
		return
	}

	// If remember is true, create refresh token
	var refreshToken string
	if loginRequest.Remember {
		refreshTokenEntity, err := c.Repository.CreateRefreshToken(tokenEntity)
		if err != nil {
			logger.Log().Err(err).Msg("Failed to create refresh token")
			api.InternalServerError(w, "Failed to create refresh token")
			return
		}

		// Encrypt refresh token
		encryptedRefreshToken, err := refreshTokenEntity.ToJWT(c.Config.SecretKey)
		if err != nil {
			logger.Log().Err(err).Msg("Failed to encrypt refresh token")
			api.InternalServerError(w, "Failed to encrypt refresh token")
			return
		}

		// Create refresh token cookie and set it in the response
		refreshToken = encryptedRefreshToken
		refreshTokenCookie := http.Cookie{
			Name:     "refreshToken",
			Value:    refreshToken,
			Path:     "/",
			MaxAge:   3600,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}

		http.SetCookie(w, &refreshTokenCookie)
	}

	json.NewEncoder(w).Encode(map[string]string{
		"access_token": encryptedToken,
	})
}

// Handles the POST request for the user logout
func (c *UserAuthController) Logout(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	logger := core.GetLogger(ctx)
	token := auth.GetCtxAccessToken(ctx)

	// Delete the access token that is currently in the context from the database
	res, err := c.Repository.DeleteAccessToken(token.ID)
	if err != nil {
		logger.Log().Err(err).Msg("Failed to delete access token")
		api.InternalServerError(w, "Failed logging out")
		return
	}

	json.NewEncoder(w).Encode(res)
}

// Handles the GET request for the user info that is currently authenticated
func (c *UserAuthController) Info(w http.ResponseWriter, req *http.Request) {
	user := auth.GetCtxUser(req.Context())

	// Return the user info as JSON
	json.NewEncoder(w).Encode(user.ToDTO())
}

// Handles the POST request for the user refresh token
func (c *UserAuthController) RefreshToken(w http.ResponseWriter, req *http.Request) {
	logger := core.GetLogger(req.Context())

	// Get the refresh token from the cookies
	refreshToken, err := req.Cookie("refreshToken")
	if err != nil {
		api.Unauthorized(w, "No refresh token provided")
		return
	}

	// Verify the refresh token
	accessTokenEntity, _, err := c.Service.VerifyRefreshToken(refreshToken.Value)
	if err != nil {
		if errors.Is(err, auth.ErrInvalidToken) {
			api.Unauthorized(w, "Invalid refresh token")
			return
		}

		if errors.Is(err, auth.ErrTokenNotFound) {
			api.Unauthorized(w, "Invalid refresh token")
			return
		}

		logger.Log().Err(err).Msg("Failed to verify refresh token")
		api.InternalServerError(w, "Failed to verify refresh token")
		return
	}

	// Refresh access token
	accessTokenEntity, err = c.Service.RefreshAccessToken(accessTokenEntity)
	if err != nil {
		logger.Log().Err(err).Msg("Failed to refresh access token")
		api.InternalServerError(w, "Failed to refresh access token")
		return
	}

	// Encrypt access token
	encryptedToken, err := accessTokenEntity.ToJWT(c.Config.SecretKey)
	if err != nil {
		logger.Log().Err(err).Msg("Failed to encrypt access token")
		api.InternalServerError(w, "Failed to encrypt access token")
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"access_token": encryptedToken,
	})
}
