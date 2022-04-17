package handlers

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func (h *Handler) usersRoutes(api *router.Group) {
	users := api.Group("/users")
	{
		users.GET("/list", h.getAllUsers)
	}
}

func (h *Handler) getAllUsers(c *fasthttp.RequestCtx) {
	c.Response.StatusCode()
}
