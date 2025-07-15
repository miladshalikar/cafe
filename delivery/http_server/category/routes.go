package categoryhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/miladshalikar/cafe/delivery/http_server/middleware"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	g := e.Group("/categories")
	g.POST("", h.AddNewCategoryHandler, middleware.Auth(h.tknSvc, h.tknCfg), middleware.Acl("create-category", h.aclSvc))
	g.DELETE("/:id", h.DeleteCategoryHandler, middleware.Auth(h.tknSvc, h.tknCfg), middleware.Acl("delete-category", h.aclSvc))
	g.GET("/:id", h.ShowCategoryHandler, middleware.Auth(h.tknSvc, h.tknCfg), middleware.Acl("index-category", h.aclSvc))
	g.PATCH("/:id", h.UpdateCategoryHandler, middleware.Auth(h.tknSvc, h.tknCfg), middleware.Acl("index-category", h.aclSvc))
	g.GET("", h.GetCategoriesHandler, middleware.Auth(h.tknSvc, h.tknCfg), middleware.Acl("index-category", h.aclSvc))
}
