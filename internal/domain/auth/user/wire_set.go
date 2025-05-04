package user

import (
	"server/internal/domain/auth/user/api"
	"server/internal/domain/auth/user/repository"
	"server/internal/domain/auth/user/service"

	"github.com/google/wire"
)

// Wires the user authentication domain
var WireSet = wire.NewSet(
	repository.NewTokenRepository,
	service.NewUserAuthService,
	api.NewUserAuthController,
	NewUserAuthDepsContainer,
)
