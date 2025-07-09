package categoryhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/miladshalikar/cafe/delivery/http_server/middleware"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	g := e.Group("/categories")
	g.POST("", h.AddNewCategory, middleware.Auth(h.tknSvc, h.tknCfg), middleware.Acl("create-category", h.aclSvc))
}
