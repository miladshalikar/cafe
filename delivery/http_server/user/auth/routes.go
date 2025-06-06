package userauthhandler

import "github.com/labstack/echo/v4"

func (h Handler) SetRoutes(e *echo.Echo) {
	g := e.Group("/users")
	g.POST("/register", h.RegisterUser)
}
