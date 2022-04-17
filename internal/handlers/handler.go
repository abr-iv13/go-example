package handlers

import (
	"project/back/internal/services"
	"project/back/pkg/logger"

	"github.com/fasthttp/router"
)

type Handler struct {
	Services *services.Services
	log      *logger.Logger
}

func NewHandler(services *services.Services, log *logger.Logger) *Handler {
	return &Handler{
		Services: services,
		log:      log,
	}
}

func (h *Handler) InitRoutes() *router.Router {
	router := router.New()

	api := router.Group("/api")
	{
		h.usersRoutes(api)
	}

	return router
}
