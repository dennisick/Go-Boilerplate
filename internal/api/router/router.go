package router

import (
	"os"
	"server/internal/api/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type Route func(r chi.Router)

// Creates a new chi router with the default middlewares
func New() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestTime)
	r.Use(middleware.RequestID)
	r.Use(middleware.ContentTypeJSON)
	r.Use(middleware.Logger(zerolog.New(os.Stdout)))

	return r
}
