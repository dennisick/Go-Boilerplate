package user

import "github.com/google/wire"

// Wires the user authentication domain
var WireSet = wire.NewSet(
	NewTokenRepository,
	NewUserAuthService,
	NewUserAuthController,
	NewUserAuthDepsContainer,
)
