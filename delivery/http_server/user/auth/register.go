package userauthhandler

import (
	"github.com/labstack/echo/v4"
	param "github.com/miladshalikar/cafe/param/authservice"
	"net/http"
)

func (h Handler) RegisterUser(ctx echo.Context) error {
	var req param.RegisterRequest
	bErr := ctx.Bind(&req)
	if bErr != nil {
		return ctx.JSON(http.StatusBadRequest, bErr)
	}

	if fieldErrors, err := h.userAuthVld.ValidateRegisterRequest(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusBadRequest, fieldErrors)
	}

	res, err := h.userAuthSvc.Register(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}
