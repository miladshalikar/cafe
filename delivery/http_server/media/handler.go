package mediahandler

import (
	liaraobjectstorage "github.com/miladshalikar/cafe/adapter/objectstorage/liara"
	aclservice "github.com/miladshalikar/cafe/service/acl"
	mediaservice "github.com/miladshalikar/cafe/service/media"
	usertokenauthservice "github.com/miladshalikar/cafe/service/user/token"
	mediavalidator "github.com/miladshalikar/cafe/validator/media"
)

type Handler struct {
	mediaSvc mediaservice.Service
	mediaVld mediavalidator.Validator
	tknSvc   usertokenauthservice.Service
	tknCfg   usertokenauthservice.Config
	aclSvc   aclservice.Service
	mediaCfg liaraobjectstorage.Config
}

func New(
	mediaSvc mediaservice.Service,
	mediaVld mediavalidator.Validator,
	tknSvc usertokenauthservice.Service,
	tknCfg usertokenauthservice.Config,
	aclSvc aclservice.Service,
	mediaCfg liaraobjectstorage.Config,
) Handler {
	return Handler{
		mediaSvc: mediaSvc,
		mediaVld: mediaVld,
		tknSvc:   tknSvc,
		tknCfg:   tknCfg,
		aclSvc:   aclSvc,
		mediaCfg: mediaCfg,
	}
}
