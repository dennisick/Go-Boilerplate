package database

import (
	"fmt"
	"log"
	"time"

	"server/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	maxOpenDatabaseConnections    = 25
	minOpenDatabaseConnections    = 25
	maxDatabaseConnectionLifetime = 5 * time.Minute
	maxDatabaseConnectionIdleTime = 5 * time.Minute
	databaseHealthCheckPeriod     = time.Minute
	databaseConnectionTimeout     = time.Second * 5
)

// Initializes a new database config
func NewConfig(config *config.DatabaseConfig) *pgxpool.Config {
	dbConfig, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database,
	))

	if err != nil {
		log.Fatal("Failed to create database pool config:", err)
	}

	dbConfig.MaxConns = maxOpenDatabaseConnections
	dbConfig.MinConns = minOpenDatabaseConnections
	dbConfig.MaxConnLifetime = maxDatabaseConnectionLifetime
	dbConfig.MaxConnIdleTime = maxDatabaseConnectionIdleTime
	dbConfig.HealthCheckPeriod = databaseHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = databaseConnectionTimeout

	return dbConfig
}
