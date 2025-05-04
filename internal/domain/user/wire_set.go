package user

import (
	"server/internal/domain/user/api"
	"server/internal/domain/user/repository"

	"github.com/google/wire"
)

// Wi
var WireSet = wire.NewSet(
	repository.NewUserRepository,
	api.NewUserController,
	NewUserDepsContainer,
)
