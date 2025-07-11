package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/miladshalikar/cafe/pkg/claims"
	aclservice "github.com/miladshalikar/cafe/service/acl"
	"net/http"
)

func Acl(requiredPermission string, aclSvc aclservice.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {

			userID, cErr := claims.GetIdFromClaim(ctx)
			if cErr != nil {
				return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
			}

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
