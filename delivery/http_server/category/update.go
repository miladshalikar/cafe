package categoryhandler

import (
	"github.com/labstack/echo/v4"
	categoryparam "github.com/miladshalikar/cafe/param/category"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	httpmsg "github.com/miladshalikar/cafe/pkg/http_message"
	"github.com/miladshalikar/cafe/pkg/logger"
	"net/http"
	"strconv"
)

func (h Handler) UpdateCategoryHandler(ctx echo.Context) error {

	id, iErr := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if iErr != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": errmsg.ErrorMsgInvalidInput,
		})
	}

	var req categoryparam.UpdateCategoryRequest

	bErr := ctx.Bind(&req)
	if bErr != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": errmsg.ErrorMsgInvalidInput,
		})
	}
	req.ID = uint(id)

	if fieldErrors, err := h.categoryVld.ValidateUpdateCategory(ctx.Request().Context(), req); err != nil {
		msg, code := httpmsg.Error(err)
		logger.Log(err)
		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	res, err := h.categorySvc.UpdateCategory(ctx.Request().Context(), req)
	if err != nil {
		msg, code := httpmsg.Error(err)
		logger.Log(err)
		return echo.NewHTTPError(code, msg)
	}

	return ctx.JSON(http.StatusOK, res)
}
