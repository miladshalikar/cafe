package userauthvalidator

import (
	"context"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	param "github.com/miladshalikar/cafe/param/user/authservice"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
	"regexp"
)

func (v Validator) ValidateLoginWithEmailRequest(ctx context.Context, req param.LoginWithEmailRequest) (map[string]string, error) {
	const op = "userauthvalidator.ValidateLoginWithEmailRequest"

	if err := validation.ValidateStructWithContext(ctx, &req,
		validation.Field(&req.Password, validation.Required, validation.NotNil),
		validation.Field(&req.Email,
			validation.Required,
			validation.Match(regexp.MustCompile(emailRegex)).Error(errmsg.ErrorMsgEmailIsNotValid),
			validation.WithContext(v.isUserExist))); err != nil {

		fieldErrors := make(map[string]string)
		vErr := validation.Errors{}
		if errors.As(err, &vErr) {
			for key, value := range vErr {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}
		return fieldErrors, richerror.New(op).WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).
			WithWarpError(err)
	}
	return nil, nil
}

func (v Validator) isUserExist(ctx context.Context, value interface{}) error {
	const op = "userauthvalidator.isUserExist"

	email, ok := value.(string)
	if !ok {
		return richerror.New(op).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindInvalid)
	}

	existed, err := v.repo.EmailExistInDB(ctx, email)
	if err != nil {
		return richerror.New(op).WithWarpError(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	if !existed {
		return richerror.New(op).WithMessage(errmsg.ErrorMsgEmailOrPassIsIncorrect).WithKind(richerror.KindNotFound)
	}
	return nil
}
