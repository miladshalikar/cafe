package categoryhandler

import (
	"github.com/labstack/echo/v4"
	categoryparam "github.com/miladshalikar/cafe/param/category"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	httpmsg "github.com/miladshalikar/cafe/pkg/http_message"
	"net/http"
)

func (h Handler) AddNewCategoryHandler(ctx echo.Context) error {

	var req categoryparam.AddNewCategoryRequest

	cErr := ctx.Bind(&req)
	if cErr != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": errmsg.ErrorMsgInvalidInput,
		})
	}

	if fieldErrors, err := h.categoryVld.ValidateAddCategory(ctx.Request().Context(), req); err != nil {
		msg, code := httpmsg.Error(err)

		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	res, err := h.categorySvc.AddNewCategory(ctx.Request().Context(), req)
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}
	return ctx.JSON(http.StatusCreated, res)
}
