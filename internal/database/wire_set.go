package database

import "github.com/google/wire"

// Wires the database domain
var WireSet = wire.NewSet(New)
