// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"project/back/config"
	"project/back/internal/repository"
	"project/back/internal/services"
	handlersHttp "project/back/internal/transport/http"
	"project/back/pkg/httpserver"
	"project/back/pkg/logger"
	"syscall"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	//Logger
	log := logger.New(cfg.Log.Level)

	//Database
	db, err := sqlx.Connect("postgres", cfg.PG.String())
	if err != nil {
		log.Fatal(err)
	}

	//Repositories
	repos := repository.NewRepositories(db)

	//Services
	services := services.NewServices(repos)

	// HTTP Server
	handlers := handlersHttp.NewHandler(services)
	httpServer := httpserver.New(handlers.InitRoutes(), httpserver.Port(cfg.HTTP.Port))

	log.Info("Starting to HTTP Server")

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
