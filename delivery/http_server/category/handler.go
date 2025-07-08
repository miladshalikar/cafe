package categoryhandler

import (
	categoryservice "github.com/miladshalikar/cafe/service/categoty"
	usertokenauthservice "github.com/miladshalikar/cafe/service/user/token"
	categoryvalidator "github.com/miladshalikar/cafe/validator/category"
)

type Handler struct {
	categorySvc categoryservice.Service
	categoryVld categoryvalidator.Validator
	tknSvc      usertokenauthservice.Service
	tknCfg      usertokenauthservice.Config
}

func New(
	categorySvc categoryservice.Service,
	categoryVld categoryvalidator.Validator,
	tknSvc usertokenauthservice.Service,
	tknCfg usertokenauthservice.Config,
) Handler {
	return Handler{
		categorySvc: categorySvc,
		categoryVld: categoryVld,
		tknSvc:      tknSvc,
		tknCfg:      tknCfg,
	}
}
