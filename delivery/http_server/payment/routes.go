package paymenthandler

import (
	"github.com/labstack/echo/v4"
	"github.com/miladshalikar/cafe/delivery/http_server/middleware"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	g := e.Group("/payments")
	g.POST("", h.CreatePaymentHandler, middleware.Auth(h.tknSvc, h.tknCfg))
	g.GET("", h.GetPaymentsHandler, middleware.Auth(h.tknSvc, h.tknCfg))
	g.GET("/:id", h.ShowPaymentHandler, middleware.Auth(h.tknSvc, h.tknCfg))
	g.PATCH("/refund/:id", h.RefundPaymentHandler, middleware.Auth(h.tknSvc, h.tknCfg))
	g.PATCH("/result/:id", h.SimulateResultHandler, middleware.Auth(h.tknSvc, h.tknCfg))
}
