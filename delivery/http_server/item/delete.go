package itemhandler

import (
	"github.com/labstack/echo/v4"
	itemparam "github.com/miladshalikar/cafe/param/item"
	"net/http"
	"strconv"
)

func (h Handler) DeleteItemHandler(ctx echo.Context) error {

	id, iErr := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if iErr != nil {
		return ctx.JSON(http.StatusBadRequest, iErr)
	}

	req := itemparam.DeleteItemRequest{ID: uint(id)}

	if fieldErrors, err := h.itemVld.ValidateDeleteItem(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, fieldErrors)
	}

	res, err := h.itemSvc.DeleteItem(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusNoContent, res)
}
