package categoryhandler

import (
	"github.com/labstack/echo/v4"
	categoryparam "github.com/miladshalikar/cafe/param/category"
	"net/http"
)

func (h Handler) AddNewCategory(ctx echo.Context) error {

	var req categoryparam.AddNewCategoryRequest

	cErr := ctx.Bind(&req)
	if cErr != nil {
		return ctx.JSON(http.StatusBadRequest, cErr)
	}

	if fieldErrors, err := h.categoryVld.ValidateAddCategory(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, fieldErrors)
	}

	res, err := h.categorySvc.AddNewCategory(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}
