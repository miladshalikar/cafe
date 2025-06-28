package userprofilehandler

import (
	userprofileservice "github.com/miladshalikar/cafe/service/user/profile"
	usertokenauthservice "github.com/miladshalikar/cafe/service/user/token"
)

type Handler struct {
	userProfSvc userprofileservice.Service
	userTknSvc  usertokenauthservice.Service
	userTknCfg  usertokenauthservice.Config
}

func New(userProfSvc userprofileservice.Service,
	userTknSvc usertokenauthservice.Service,
	userTknCfg usertokenauthservice.Config) Handler {
	return Handler{userProfSvc: userProfSvc, userTknSvc: userTknSvc, userTknCfg: userTknCfg}
}
