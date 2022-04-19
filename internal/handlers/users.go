package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func (h *Handler) usersRoutes(api *gin.RouterGroup) {
// 	users := api.Group("/users")
// 	{
// 		users.GET("/list", h.getAllUsers)
// 	}
// }

func (h *Handler) getAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, "hello")
}
