package connection

import (
	"server/internal/domain/connection/api"
	"server/internal/domain/connection/repository"

	"github.com/google/wire"
)

// Wires the connection domain
var WireSet = wire.NewSet(
	repository.NewConnectionRepository,
	api.NewConnectionController,
	NewConnectionDepsContainer,
)
