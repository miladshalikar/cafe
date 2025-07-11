package userprofilehandler

import (
	"github.com/labstack/echo/v4"
	userprofileserviceparam "github.com/miladshalikar/cafe/param/user/profile"
	claims "github.com/miladshalikar/cafe/pkg/claims"
	"net/http"
)

func (h Handler) Profile(ctx echo.Context) error {

	var req userprofileserviceparam.UserProfileRequest

	id, cErr := claims.GetIdFromClaim(ctx)
	if cErr != nil {
		return ctx.JSON(401, map[string]string{"error": "unauthorized"})
	}

	req.Id = int(id)

	res, err := h.userProfSvc.GetUserByID(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "something wrong")
	}
	return ctx.JSON(200, res)
}
