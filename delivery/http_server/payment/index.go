package paymenthandler

import (
	"github.com/labstack/echo/v4"
	commonparam "github.com/miladshalikar/cafe/param/common"
	paymentparam "github.com/miladshalikar/cafe/param/payment"
	httpmsg "github.com/miladshalikar/cafe/pkg/http_message"
	"github.com/miladshalikar/cafe/pkg/logger"
	"net/http"
	"strconv"
)

func (h Handler) GetPaymentsHandler(ctx echo.Context) error {

	pageNumber, nErr := strconv.ParseUint(ctx.QueryParam("page_number"), 10, 64)
	if nErr != nil {
		pageNumber = commonparam.DefaultPageNumber
	}

	pageSize, sErr := strconv.ParseUint(ctx.QueryParam("page_size"), 10, 64)
	if sErr != nil {
		pageSize = commonparam.DefaultPageSize
	}

	OrderID, cErr := strconv.ParseUint(ctx.QueryParam("order_id"), 10, 64)
	if cErr != nil {
		OrderID = 0
	}

	req := paymentparam.GetPaymentsByOrderIDRequest{
		Pagination: commonparam.PaginationRequest{
			PageSize:   uint(pageSize),
			PageNumber: uint(pageNumber),
		},
		OrderID: uint(OrderID),
	}

	res, err := h.paymentSvc.GetPaymentsByOrderID(ctx.Request().Context(), req)
	if err != nil {
		msg, code := httpmsg.Error(err)
		logger.Log(err)
		return echo.NewHTTPError(code, msg)
	}

	return ctx.JSON(http.StatusOK, res)
}
