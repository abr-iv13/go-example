package handlers

import (
	"project/back/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		middleware.CorsMiddleware(),
	)

	api := router.Group("/api")
	{
		// h.usersRoutes(api)
		users := api.Group("/users")
		{
			users.GET("/list", h.getAllUsers)
		}
	}

	return router
}
