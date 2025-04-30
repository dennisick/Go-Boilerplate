package api

import (
	"fmt"
	"net/http"
	"server/internal/config"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
)

const (
	httpReadTimeout  = 3 * time.Second // The maximum time to read the request
	httpWriteTimeout = 5 * time.Second // The maximum time to write the response
	httpIdleTimeout  = 5 * time.Second // The maximum time to idle
)

var server *http.Server
var once sync.Once

// Creates a new HTTP server with the given configuration and router
// It is a singleton, so it will return the same server for each call
func New(config *config.ApplicationConfig, router chi.Router) *http.Server {
	once.Do(func() {
		server = &http.Server{
			Addr:         fmt.Sprintf(":%d", config.Http.Port),
			Handler:      router,
			ReadTimeout:  httpReadTimeout,
			WriteTimeout: httpWriteTimeout,
			IdleTimeout:  httpIdleTimeout,
		}
	})

	return server
}
