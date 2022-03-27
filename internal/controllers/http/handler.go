package http

import (
	"project/back/internal/services"
	"project/back/pkg/logger"
	"project/back/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services *services.Services
	log      *logger.Logger
}

func NewHandler(services *services.Services, log *logger.Logger) *Handler {
	return &Handler{Services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		middleware.CorsMiddleware(),
	)

	api := router.Group("/api")
	{
		h.addUsersRoutes(api)
	}

	return router
}
