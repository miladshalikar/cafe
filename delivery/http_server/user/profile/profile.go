package userprofilehandler

import (
	"github.com/labstack/echo/v4"
	userprofileserviceparam "github.com/miladshalikar/cafe/param/user/profile"
	claims "github.com/miladshalikar/cafe/pkg/claims"
	httpmsg "github.com/miladshalikar/cafe/pkg/http_message"
	"github.com/miladshalikar/cafe/pkg/logger"
	"net/http"
)

func (h Handler) Profile(ctx echo.Context) error {

	var req userprofileserviceparam.UserProfileRequest

	id, cErr := claims.GetIdFromClaim(ctx)
	if cErr != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"message": "unauthorized"})
	}

	req.Id = int(id)

	res, err := h.userProfSvc.GetUserByID(ctx.Request().Context(), req)
	if err != nil {
		msg, code := httpmsg.Error(err)
		logger.Log(err)
		return echo.NewHTTPError(code, msg)
	}
	return ctx.JSON(http.StatusOK, res)
}
