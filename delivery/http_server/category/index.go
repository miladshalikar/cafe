package categoryhandler

import (
	"github.com/labstack/echo/v4"
	categoryparam "github.com/miladshalikar/cafe/param/category"
	commonparam "github.com/miladshalikar/cafe/param/common"
	"net/http"
	"strconv"
)

func (h Handler) GetCategoriesHandler(ctx echo.Context) error {

	pageNumber, nErr := strconv.ParseUint(ctx.QueryParam("page_number"), 10, 64)
	if nErr != nil {
		pageNumber = commonparam.DefaultPageNumber
	}

	pageSize, sErr := strconv.ParseUint(ctx.QueryParam("page_size"), 10, 64)
	if sErr != nil {
		pageSize = commonparam.DefaultPageSize
	}

	search := ctx.QueryParam("name")

	req := categoryparam.GetCategoriesRequest{
		Pagination: commonparam.PaginationRequest{
			PageSize:   uint(pageSize),
			PageNumber: uint(pageNumber),
		},
		Search: commonparam.SearchRequest{Search: search},
	}

	res, err := h.categorySvc.GetCategories(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, res)
}
