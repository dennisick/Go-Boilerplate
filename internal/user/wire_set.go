package user

import (
	"github.com/google/wire"
)

// Wi
var WireSet = wire.NewSet(
	NewUserRepository,
	NewUserController,
	NewUserDepsContainer,
)
