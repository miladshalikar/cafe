package userauthhandler

import (
	"github.com/miladshalikar/cafe/service/user/authservice"
	userauthvalidator "github.com/miladshalikar/cafe/validator/user/auth"
)

type Handler struct {
	userAuthSvc userauthservice.Service
	userAuthVld userauthvalidator.Validator
}

func New(userAuthSvc userauthservice.Service, userAuthVld userauthvalidator.Validator) Handler {
	return Handler{userAuthSvc: userAuthSvc, userAuthVld: userAuthVld}
}
