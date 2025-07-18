package mediahandler

import (
	"github.com/labstack/echo/v4"
	mediaparam "github.com/miladshalikar/cafe/param/media"
	"net/http"
	"strconv"
)

func (h Handler) GetURL(ctx echo.Context) error {

	num, cErr := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if cErr != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid id",
		})
	}

	if fieldErrors, err := h.mediaVld.ValidateGetFile(ctx.Request().Context(), mediaparam.GetMediaRequest{ID: uint(num)}); err != nil {
		return ctx.JSON(http.StatusBadRequest, fieldErrors)
	}

	resp, sErr := h.mediaSvc.GetURLMedia(ctx.Request().Context(), mediaparam.GetURLRequest{
		ID: uint(num),
	})
	if sErr != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": sErr.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, resp)
}
