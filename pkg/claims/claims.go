package claims

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/miladshalikar/cafe/config"
	usertokenauthservice "github.com/miladshalikar/cafe/service/user/token"
)

func GetIdFromClaim(ctx echo.Context) (uint, error) {
	claims := ctx.Get(config.AuthMiddlewareContextKey).(*usertokenauthservice.Claims)
	if claims == nil {
		return 0, errors.New("claims not found")
	}
	userID := claims.UserID

	return userID, nil
}
