package main

import (
	"context"
	"log"
)

func main() {
	// Init app
	app := InitializeApp()

	// Test database connection
	defer app.Database.Close()

	testConnection, err := app.Database.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error while acquiring connection from the database pool:", err)
	}

	err = testConnection.Ping(context.Background())
	if err != nil {
		log.Fatal("Error while pinging database:", err)
	}
	testConnection.Release()

	log.Println("Database successfully connected")

	log.Printf("HTTP server listening...")
	app.Api.ListenAndServe()
}
