package http

import (
	"project/back/internal/services"
	"project/back/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services *services.Services
}

func NewHandler(services *services.Services) *Handler {
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
