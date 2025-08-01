package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/miladshalikar/cafe/pkg/claims"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	aclservice "github.com/miladshalikar/cafe/service/acl"
	"net/http"
)

func Acl(requiredPermission string, aclSvc aclservice.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {

			userID, cErr := claims.GetIdFromClaim(ctx)
			if cErr != nil {
				return ctx.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
			}

			hasPermission, err := aclSvc.HasPermission(userID, requiredPermission)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, errmsg.ErrorMsgSomethingWentWrong)
			}
			if !hasPermission {
				return ctx.JSON(http.StatusForbidden, echo.Map{"message": "access denied"})
			}

			return next(ctx)
		}
	}
}
