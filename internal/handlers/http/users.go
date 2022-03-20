package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.GET("/", h.getAllUsers)
	}
}

func (h *Handler) getAllUsers(c *gin.Context) {

	c.JSON(http.StatusOK, "hello")
}
