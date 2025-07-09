package categoryhandler

import (
	aclservice "github.com/miladshalikar/cafe/service/acl"
	categoryservice "github.com/miladshalikar/cafe/service/categoty"
	usertokenauthservice "github.com/miladshalikar/cafe/service/user/token"
	categoryvalidator "github.com/miladshalikar/cafe/validator/category"
)

type Handler struct {
	categorySvc categoryservice.Service
	categoryVld categoryvalidator.Validator
	tknSvc      usertokenauthservice.Service
	tknCfg      usertokenauthservice.Config
	aclSvc      aclservice.Service
}

func New(
	categorySvc categoryservice.Service,
	categoryVld categoryvalidator.Validator,
	tknSvc usertokenauthservice.Service,
	tknCfg usertokenauthservice.Config,
	aclSvc aclservice.Service,
) Handler {
	return Handler{
		categorySvc: categorySvc,
		categoryVld: categoryVld,
		tknSvc:      tknSvc,
		tknCfg:      tknCfg,
		aclSvc:      aclSvc,
	}
}
