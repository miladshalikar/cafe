package itemhandler

import (
	"github.com/labstack/echo/v4"
	itemparam "github.com/miladshalikar/cafe/param/item"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	httpmsg "github.com/miladshalikar/cafe/pkg/http_message"
	"github.com/miladshalikar/cafe/pkg/logger"
	"net/http"
)

func (h Handler) AddNewItemHandler(ctx echo.Context) error {

	var req itemparam.AddNewItemRequest

	cErr := ctx.Bind(&req)
	if cErr != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": errmsg.ErrorMsgInvalidInput,
		})
	}

	if fieldErrors, err := h.itemVld.ValidateAddItem(ctx.Request().Context(), req); err != nil {
		msg, code := httpmsg.Error(err)
		logger.Log(err)
		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	res, err := h.itemSvc.AddNewItem(ctx.Request().Context(), req)
	if err != nil {
		msg, code := httpmsg.Error(err)
		logger.Log(err)
		return echo.NewHTTPError(code, msg)
	}
	return ctx.JSON(http.StatusOK, res)
}
