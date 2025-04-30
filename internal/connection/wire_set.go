package connection

import "github.com/google/wire"

// Wires the connection domain
var WireSet = wire.NewSet(
	NewConnectionRepository,
	NewConnectionController,
	NewConnectionDepsContainer,
)
