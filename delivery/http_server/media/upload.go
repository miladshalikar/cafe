package mediahandler

import (
	"github.com/labstack/echo/v4"
	mediaparam "github.com/miladshalikar/cafe/param/media"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	httpmsg "github.com/miladshalikar/cafe/pkg/http_message"
	"net/http"
	"strconv"
)

func (h Handler) UploadMedia(ctx echo.Context) error {

	file, fErr := ctx.FormFile("file")
	if fErr != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": errmsg.ErrorMsgInvalidInput,
		})
	}

	isPrivateStr := ctx.FormValue("is_private")
	isPrivate, bErr := strconv.ParseBool(isPrivateStr)
	if bErr != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": errmsg.ErrorMsgInvalidInput,
		})
	}

	req := mediaparam.UploadMediaRequest{
		FileHeader: file,
		IsPrivate:  isPrivate,
		Bucket:     h.mediaCfg.BucketName,
	}

	if fieldErrors, vErr := h.mediaVld.ValidateUploadFile(ctx.Request().Context(), req); vErr != nil {
		msg, code := httpmsg.Error(vErr)

		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	res, mErr := h.mediaSvc.UploadMedia(ctx.Request().Context(), req)
	if mErr != nil {
		msg, code := httpmsg.Error(mErr)
		return echo.NewHTTPError(code, msg)
	}
	return ctx.JSON(http.StatusOK, res)

}
