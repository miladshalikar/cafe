package userprofilehandler

import (
	"github.com/labstack/echo/v4"
	"github.com/miladshalikar/cafe/delivery/http_server/middleware"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	g := e.Group("/users")
	g.GET("/profile", h.Profile, middleware.Auth(h.userTknSvc, h.userTknCfg))
}
