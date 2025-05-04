// wire.go
//go:build wireinject

package main

import (
	"net/http"
	"server/internal/api"
	"server/internal/api/router"
	"server/internal/config"
	"server/internal/database"
	"server/internal/domain/application"
	userAuth "server/internal/domain/auth/user"
	"server/internal/domain/connection"
	"server/internal/domain/user"

	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Config          *config.ApplicationConfig
	UserDeps        *user.UserDepsContainer
	ApplicationDeps *application.ApplicationDepsContainer
	ConnectionDeps  *connection.ConnectionDepsContainer
	UserAuthDeps    *userAuth.UserAuthDepsContainer
	Database        *pgxpool.Pool
	Api             *http.Server
}

func NewApp(
	config *config.ApplicationConfig,
	userDeps *user.UserDepsContainer,
	applicationDeps *application.ApplicationDepsContainer,
	connectionDeps *connection.ConnectionDepsContainer,
	userAuthDeps *userAuth.UserAuthDepsContainer,
	db *pgxpool.Pool,
) *App {
	r := router.New()
	r.Group(userDeps.RegisterRoutes(r))
	r.Group(applicationDeps.RegisterRoutes(r))
	r.Group(connectionDeps.RegisterRoutes(r))
	r.Group(userAuthDeps.RegisterRoutes(r))

	api := api.New(config, r)

	return &App{
		Config:          config,
		UserDeps:        userDeps,
		ApplicationDeps: applicationDeps,
		ConnectionDeps:  connectionDeps,
		UserAuthDeps:    userAuthDeps,
		Database:        db,
		Api:             api,
	}
}

func InitializeApp() *App {
	wire.Build(
		config.LoadConfig,
		database.WireSet,
		user.WireSet,
		application.WireSet,
		connection.WireSet,
		userAuth.WireSet,
		NewApp,
	)

	return &App{}
}
