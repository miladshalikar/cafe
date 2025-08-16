package userauthhandler

import (
	"github.com/labstack/echo/v4"
	param "github.com/miladshalikar/cafe/param/user/authservice"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	httpmsg "github.com/miladshalikar/cafe/pkg/http_message"
	"github.com/miladshalikar/cafe/pkg/logger"
	"net/http"
)

func (h Handler) RegisterUser(ctx echo.Context) error {
	var req param.RegisterRequest
	bErr := ctx.Bind(&req)
	if bErr != nil {
		return ctx.JSON(http.StatusBadRequest, errmsg.ErrorMsgInvalidInput)
	}

	if fieldErrors, err := h.userAuthVld.ValidateRegisterRequest(ctx.Request().Context(), req); err != nil {
		msg, code := httpmsg.Error(err)
		logger.Log(err)
		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	res, err := h.userAuthSvc.Register(ctx.Request().Context(), req)
	if err != nil {
		msg, code := httpmsg.Error(err)
		logger.Log(err)
		return echo.NewHTTPError(code, msg)
	}
	return ctx.JSON(http.StatusOK, res)
}
