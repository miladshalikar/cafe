package itemhandler

import (
	aclservice "github.com/miladshalikar/cafe/service/acl"
	itemservice "github.com/miladshalikar/cafe/service/item"
	usertokenauthservice "github.com/miladshalikar/cafe/service/user/token"
	itemvalidator "github.com/miladshalikar/cafe/validator/item"
)

type Handler struct {
	itemSvc itemservice.Service
	itemVld itemvalidator.Validator
	tknSvc  usertokenauthservice.Service
	tknCfg  usertokenauthservice.Config
	aclSvc  aclservice.Service
}

func New(
	itemSvc itemservice.Service,
	itemVld itemvalidator.Validator,
	tknSvc usertokenauthservice.Service,
	tknCfg usertokenauthservice.Config,
	aclSvc aclservice.Service,
) Handler {
	return Handler{
		itemSvc: itemSvc,
		itemVld: itemVld,
		tknSvc:  tknSvc,
		tknCfg:  tknCfg,
		aclSvc:  aclSvc,
	}
}
