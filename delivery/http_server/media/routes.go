package mediahandler

import (
	"github.com/labstack/echo/v4"
	"github.com/miladshalikar/cafe/delivery/http_server/middleware"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	g := e.Group("/media")
	g.POST("/upload", h.UploadMedia, middleware.Auth(h.tknSvc, h.tknCfg))
	g.GET("/get-url/:id", h.GetURL, middleware.Auth(h.tknSvc, h.tknCfg))
	g.DELETE("/delete/:id", h.DeleteMedia, middleware.Auth(h.tknSvc, h.tknCfg))
}
