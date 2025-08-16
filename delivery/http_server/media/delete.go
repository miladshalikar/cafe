package mediahandler

import (
	"github.com/labstack/echo/v4"
	mediaparam "github.com/miladshalikar/cafe/param/media"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	httpmsg "github.com/miladshalikar/cafe/pkg/http_message"
	"github.com/miladshalikar/cafe/pkg/logger"
	"net/http"
	"strconv"
)

func (h Handler) DeleteMedia(ctx echo.Context) error {

	num, cErr := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if cErr != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": errmsg.ErrorMsgInvalidInput,
		})
	}

	if fieldErrors, err := h.mediaVld.ValidateGetFile(ctx.Request().Context(), mediaparam.GetMediaRequest{ID: uint(num)}); err != nil {
		msg, code := httpmsg.Error(err)
		logger.Log(err)
		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	resp, sErr := h.mediaSvc.DeleteMedia(ctx.Request().Context(), mediaparam.DeleteMediaRequest{
		ID: uint(num),
	})
	if sErr != nil {
		msg, code := httpmsg.Error(sErr)
		logger.Log(sErr)
		return echo.NewHTTPError(code, msg)
	}

	return ctx.JSON(http.StatusNoContent, resp)
}
