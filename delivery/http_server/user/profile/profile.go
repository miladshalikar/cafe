package userprofilehandler

import (
	"github.com/labstack/echo/v4"
	"github.com/miladshalikar/cafe/config"
	userprofileserviceparam "github.com/miladshalikar/cafe/param/user/profile"
	usertokenauthservice "github.com/miladshalikar/cafe/service/user/token"
	"net/http"
)

func (h Handler) Profile(ctx echo.Context) error {

	var req userprofileserviceparam.UserProfileRequest

	claims := ctx.Get(config.AuthMiddlewareContextKey).(*usertokenauthservice.Claims)
	if claims == nil {
		return ctx.JSON(401, map[string]string{"error": "unauthorized"})
	}
	req.Id = int(claims.UserID)

	res, err := h.userProfSvc.GetUserByID(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "something wrong")
	}
	return ctx.JSON(200, res)
}
