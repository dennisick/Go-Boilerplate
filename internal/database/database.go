package database

import (
	"context"
	"log"
	"sync"

	"server/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var once sync.Once
var pool *pgxpool.Pool

// Initializes a new database pool that only exists once
func New(config *config.ApplicationConfig) *pgxpool.Pool {
	once.Do(func() {
		newPool, err := pgxpool.NewWithConfig(
			context.Background(),
			NewConfig(config.Database),
		)

		if err != nil {
			log.Fatal("Error while creating connection to database:", err)
		}

		pool = newPool
	})

	return pool
}
