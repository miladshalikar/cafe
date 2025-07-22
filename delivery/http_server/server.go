package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/miladshalikar/cafe/config"
	categoryhandler "github.com/miladshalikar/cafe/delivery/http_server/category"
	itemhandler "github.com/miladshalikar/cafe/delivery/http_server/item"
	mediahandler "github.com/miladshalikar/cafe/delivery/http_server/media"
	userauthhandler "github.com/miladshalikar/cafe/delivery/http_server/user/auth"
	userprofilehandler "github.com/miladshalikar/cafe/delivery/http_server/user/profile"
)

type Server struct {
	config             config.Config
	Router             *echo.Echo
	userAuthHandler    userauthhandler.Handler
	userProfileHandler userprofilehandler.Handler
	categoryHandler    categoryhandler.Handler
	mediaHandler       mediahandler.Handler
	itemHandler        itemhandler.Handler
}

func New(config config.Config,
	userAuthHandler userauthhandler.Handler,
	userProfileHandler userprofilehandler.Handler,
	categoryHandler categoryhandler.Handler,
	mediaHandler mediahandler.Handler,
	itemHandler itemhandler.Handler,
) Server {
	return Server{
		config:             config,
		Router:             echo.New(),
		userAuthHandler:    userAuthHandler,
		userProfileHandler: userProfileHandler,
		categoryHandler:    categoryHandler,
		mediaHandler:       mediaHandler,
		itemHandler:        itemHandler,
	}
}

func (s Server) Serve() {
	s.userAuthHandler.SetRoutes(s.Router)
	s.userProfileHandler.SetRoutes(s.Router)
	s.categoryHandler.SetRoutes(s.Router)
	s.mediaHandler.SetRoutes(s.Router)
	s.itemHandler.SetRoutes(s.Router)

	address := fmt.Sprintf(":%d", s.config.Server.Port)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Router.Start(address); err != nil {
		fmt.Println("router start error", err)
	}
}
