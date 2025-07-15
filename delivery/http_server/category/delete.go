package categoryhandler

import (
	"github.com/labstack/echo/v4"
	categoryparam "github.com/miladshalikar/cafe/param/category"
	"net/http"
	"strconv"
)

func (h Handler) DeleteCategoryHandler(ctx echo.Context) error {

	id, iErr := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if iErr != nil {
		return ctx.JSON(http.StatusBadRequest, iErr)
	}

	req := categoryparam.DeleteCategoryRequest{ID: uint(id)}

	if fieldErrors, err := h.categoryVld.ValidateDeleteCategory(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, fieldErrors)
	}

	res, err := h.categorySvc.DeleteCategory(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusNoContent, res)
}
