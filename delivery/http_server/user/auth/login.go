package userauthhandler

import (
	"github.com/labstack/echo/v4"
	param "github.com/miladshalikar/cafe/param/user/authservice"
	"net/http"
)

func (h Handler) LoginUserByEmail(ctx echo.Context) error {
	var req param.LoginWithEmailRequest
	bErr := ctx.Bind(&req)
	if bErr != nil {
		return ctx.JSON(http.StatusBadRequest, bErr)
	}

	if fieldErrors, err := h.userAuthVld.ValidateLoginWithEmailRequest(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusBadRequest, fieldErrors)
	}

	res, err := h.userAuthSvc.LoginWithEmail(ctx.Request().Context(), req)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "wrong password")
	}
	return ctx.JSON(http.StatusOK, res)
}
