package paymenthandler

import (
	"github.com/labstack/echo/v4"
	paymentparam "github.com/miladshalikar/cafe/param/payment"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	httpmsg "github.com/miladshalikar/cafe/pkg/http_message"
	"github.com/miladshalikar/cafe/pkg/logger"
	"net/http"
	"strconv"
)

func (h Handler) SimulateResultHandler(ctx echo.Context) error {

	id, iErr := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if iErr != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": errmsg.ErrorMsgInvalidInput,
		})
	}

	SuccessStr := ctx.FormValue("success")
	Success, pErr := strconv.ParseBool(SuccessStr)
	if pErr != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": errmsg.ErrorMsgInvalidInput,
		})
	}

	req := paymentparam.SimulatePaymentResultRequest{
		PaymentID: uint(id),
		Success:   Success,
	}

	if fieldErrors, err := h.paymentVld.ValidateSimulateResultPayment(ctx.Request().Context(), req); err != nil {
		msg, code := httpmsg.Error(err)
		logger.Log(err)
		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	res, err := h.paymentSvc.SimulatePaymentResult(ctx.Request().Context(), req)
	if err != nil {
		msg, code := httpmsg.Error(err)
		logger.Log(err)
		return echo.NewHTTPError(code, msg)
	}

	return ctx.JSON(http.StatusNoContent, res)
}
