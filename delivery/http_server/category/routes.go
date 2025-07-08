package categoryhandler

import "github.com/labstack/echo/v4"

func (h Handler) SetRoutes(e *echo.Echo) {
	g := e.Group("/categories")
	g.POST("")
}
