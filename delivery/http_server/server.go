package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/miladshalikar/cafe/config"
	userauthhandler "github.com/miladshalikar/cafe/delivery/http_server/user/auth"
	userprofilehandler "github.com/miladshalikar/cafe/delivery/http_server/user/profile"
)

type Server struct {
	config             config.Config
	Router             *echo.Echo
	userAuthHandler    userauthhandler.Handler
	userProfileHandler userprofilehandler.Handler
}

func New(config config.Config,
	userAuthHandler userauthhandler.Handler,
	userProfileHandler userprofilehandler.Handler) Server {
	return Server{
		config:             config,
		Router:             echo.New(),
		userAuthHandler:    userAuthHandler,
		userProfileHandler: userProfileHandler,
	}
}

func (s Server) Serve() {
	s.userAuthHandler.SetRoutes(s.Router)
	s.userProfileHandler.SetRoutes(s.Router)

	address := fmt.Sprintf(":%d", s.config.Server.Port)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Router.Start(address); err != nil {
		fmt.Println("router start error", err)
	}
}
