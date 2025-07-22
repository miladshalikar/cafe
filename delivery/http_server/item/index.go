package itemhandler

import (
	"github.com/labstack/echo/v4"
	commonparam "github.com/miladshalikar/cafe/param/common"
	itemparam "github.com/miladshalikar/cafe/param/item"
	"net/http"
	"strconv"
)

func (h Handler) GetItemsHandler(ctx echo.Context) error {

	pageNumber, nErr := strconv.ParseUint(ctx.QueryParam("page_number"), 10, 64)
	if nErr != nil {
		pageNumber = commonparam.DefaultPageNumber
	}

	pageSize, sErr := strconv.ParseUint(ctx.QueryParam("page_size"), 10, 64)
	if sErr != nil {
		pageSize = commonparam.DefaultPageSize
	}

	search := ctx.QueryParam("name")

	req := itemparam.GetItemsRequest{
		Pagination: commonparam.PaginationRequest{
			PageSize:   uint(pageSize),
			PageNumber: uint(pageNumber),
		},
		Search: commonparam.SearchRequest{Search: search},
	}

	res, err := h.itemSvc.GetItems(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, res)
}
