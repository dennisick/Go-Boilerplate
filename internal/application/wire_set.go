package application

import "github.com/google/wire"

// Wires the application domain
var WireSet = wire.NewSet(
	NewApplicationRepository,
	NewApplicationController,
	NewApplicationDepsContainer,
)
