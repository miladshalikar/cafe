package paymenthandler

import (
	paymentservice "github.com/miladshalikar/cafe/service/payment"
	usertokenauthservice "github.com/miladshalikar/cafe/service/user/token"
	paymentvalidator "github.com/miladshalikar/cafe/validator/payment"
)

type Handler struct {
	paymentSvc paymentservice.Service
	paymentVld paymentvalidator.Validator
	tknSvc     usertokenauthservice.Service
	tknCfg     usertokenauthservice.Config
}

func New(
	paymentSvc paymentservice.Service,
	paymentVld paymentvalidator.Validator,
	tknSvc usertokenauthservice.Service,
	tknCfg usertokenauthservice.Config,
) Handler {
	return Handler{
		paymentSvc: paymentSvc,
		paymentVld: paymentVld,
		tknSvc:     tknSvc,
		tknCfg:     tknCfg,
	}
}
