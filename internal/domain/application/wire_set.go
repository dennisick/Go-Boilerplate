package application

import (
	"server/internal/domain/application/api"
	"server/internal/domain/application/repository"

	"github.com/google/wire"
)

// Wires the application domain
var WireSet = wire.NewSet(
	repository.NewApplicationRepository,
	api.NewApplicationController,
	NewApplicationDepsContainer,
)
