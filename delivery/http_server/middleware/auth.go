package middleware

import (
	mw "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/miladshalikar/cafe/config"
	usertokenauthservice "github.com/miladshalikar/cafe/service/user/token"
)

func Auth(service usertokenauthservice.Service, cfg usertokenauthservice.Config) echo.MiddlewareFunc {
	return mw.WithConfig(mw.Config{
		ContextKey:    config.AuthMiddlewareContextKey,
		SigningKey:    []byte(cfg.SignKey),
		SigningMethod: "HS256",
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			claims, err := service.ParseToken(auth)
			if err != nil {
				return nil, err
			}
			return claims, nil
		},
	})
}
