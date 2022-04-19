// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"project/back/config"
	"project/back/internal/handlers"
	"project/back/internal/repository"
	"project/back/internal/services"
	httpserver "project/back/pkg/http-server"
	"project/back/pkg/logger"
	"project/back/pkg/postgres"
	"syscall"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	//Logger
	log := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	//Repositories
	repos := repository.NewRepository(pg)

	//Services
	services := services.NewServices(repos)

	// HTTP Server
	handlers := handlers.NewHandler(services, log)
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
