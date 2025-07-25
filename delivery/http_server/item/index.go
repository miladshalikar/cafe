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

	CategoryID, cErr := strconv.ParseUint(ctx.QueryParam("category_id"), 10, 64)
	if cErr != nil {
		CategoryID = 0
	}

	minPrice, minErr := strconv.ParseUint(ctx.QueryParam("min_price"), 10, 64)
	if minErr != nil {
		minPrice = 0
	}

	maxPrice, maxErr := strconv.ParseUint(ctx.QueryParam("max_price"), 10, 64)
	if maxErr != nil {
		maxPrice = 0
	}

	req := itemparam.GetItemsRequest{
		Pagination: commonparam.PaginationRequest{
			PageSize:   uint(pageSize),
			PageNumber: uint(pageNumber),
		},
		Search: commonparam.SearchRequest{Search: search},
		Filter: itemparam.FilterRequest{
			CategoryID: uint(CategoryID),
			MinPrice:   uint(minPrice),
			MaxPrice:   uint(maxPrice),
		},
	}

	res, err := h.itemSvc.GetItems(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, res)
}
