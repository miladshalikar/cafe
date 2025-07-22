package itemhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/miladshalikar/cafe/delivery/http_server/middleware"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	g := e.Group("/items")
	g.POST("", h.AddNewItemHandler, middleware.Auth(h.tknSvc, h.tknCfg), middleware.Acl("create-item", h.aclSvc))
	g.DELETE("/:id", h.DeleteItemHandler, middleware.Auth(h.tknSvc, h.tknCfg), middleware.Acl("delete-item", h.aclSvc))
	g.GET("/:id", h.ShowItemHandler, middleware.Auth(h.tknSvc, h.tknCfg), middleware.Acl("index-item", h.aclSvc))
	g.PATCH("/:id", h.UpdateItemHandler, middleware.Auth(h.tknSvc, h.tknCfg), middleware.Acl("edit-item", h.aclSvc))
	g.GET("", h.GetItemsHandler, middleware.Auth(h.tknSvc, h.tknCfg), middleware.Acl("index-item", h.aclSvc))
}
