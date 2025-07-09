package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/miladshalikar/cafe/config"
	aclservice "github.com/miladshalikar/cafe/service/acl"
	usertokenauthservice "github.com/miladshalikar/cafe/service/user/token"
	"net/http"
)

func Acl(requiredPermission string, aclSvc aclservice.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {

			claims := ctx.Get(config.AuthMiddlewareContextKey).(*usertokenauthservice.Claims)
			if claims == nil {
				return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
			}

			userID := claims.UserID

			hasPermission, err := aclSvc.HasPermission(userID, requiredPermission)
			if err != nil {
				fmt.Println(err)
				return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "permission check error"})
			}
			if !hasPermission {
				return ctx.JSON(http.StatusForbidden, echo.Map{"error": "access denied"})
			}

			return next(ctx)
		}
	}
}
