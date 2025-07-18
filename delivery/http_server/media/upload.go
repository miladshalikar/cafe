package mediahandler

import (
	"github.com/labstack/echo/v4"
	mediaparam "github.com/miladshalikar/cafe/param/media"
	"net/http"
	"strconv"
)

func (h Handler) UploadMedia(ctx echo.Context) error {

	file, fErr := ctx.FormFile("file")
	if fErr != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "file is required!"})
	}

	isPrivateStr := ctx.FormValue("is_private")
	isPrivate, err := strconv.ParseBool(isPrivateStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "is_private must be a valid boolean"})
	}

	bucket := ctx.FormValue("bucket")
	if bucket == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "bucket is required"})
	}

	req := mediaparam.UploadMediaRequest{
		FileHeader: file,
		IsPrivate:  isPrivate,
		Bucket:     bucket,
	}

	if fieldErrors, err := h.mediaVld.ValidateUploadFile(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusBadRequest, fieldErrors)
	}

	res, err := h.mediaSvc.UploadMedia(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)

}
