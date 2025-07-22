package itemhandler

import (
	"github.com/labstack/echo/v4"
	itemparam "github.com/miladshalikar/cafe/param/item"
	"net/http"
)

func (h Handler) AddNewItemHandler(ctx echo.Context) error {

	var req itemparam.AddNewItemRequest

	cErr := ctx.Bind(&req)
	if cErr != nil {
		return ctx.JSON(http.StatusBadRequest, cErr)
	}

	if fieldErrors, err := h.itemVld.ValidateAddItem(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, fieldErrors)
	}

	res, err := h.itemSvc.AddNewItem(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}
