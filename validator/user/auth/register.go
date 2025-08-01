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

func (v Validator) ValidateRegisterRequest(ctx context.Context, req param.RegisterRequest) (map[string]string, error) {
	const op = "userauthvalidator.ValidateRegisterRequest"

	if err := validation.ValidateStructWithContext(ctx, &req,
		validation.Field(&req.FirstName, validation.Required, validation.Length(2, 50)),
		validation.Field(&req.LastName, validation.Required, validation.Length(2, 50)),
		validation.Field(&req.Email,
			validation.Required,
			validation.Length(5, 100),
			validation.Match(regexp.MustCompile(emailRegex)).Error(errmsg.ErrorMsgEmailIsNotValid),
			validation.WithContext(v.isEmailExistInDB)),
		validation.Field(&req.PhoneNumber,
			validation.Required,
			validation.Length(5, 100),
			validation.Match(regexp.MustCompile(phoneNumberRegex)).Error(errmsg.ErrorMsgPhoneNumberIsNotValid),
			validation.WithContext(v.isPhoneNumberExistInDB)),
		validation.Field(&req.Password, validation.Required, validation.Length(6, 100))); err != nil {

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

func (v Validator) isEmailExistInDB(ctx context.Context, value interface{}) error {
	const op = "userauthvalidator.isEmailExistInDB"

	email, ok := value.(string)
	if !ok {
		return richerror.New(op).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindInvalid)
	}
	existed, err := v.repo.EmailExistInDB(ctx, email)
	if err != nil {
		return richerror.New(op).WithWarpError(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	if existed {
		return richerror.New(op).WithMessage(errmsg.ErrorMsgEmailIsNotUnique)
	}
	return nil
}

func (v Validator) isPhoneNumberExistInDB(ctx context.Context, value interface{}) error {
	const op = "userauthvalidator.isPhoneNumberExistInDB"

	phoneNumber, ok := value.(string)
	if !ok {
		return richerror.New(op).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindInvalid)
	}
	existed, err := v.repo.PhoneNumberExistInDB(ctx, phoneNumber)
	if err != nil {
		return richerror.New(op).WithWarpError(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	if existed {
		return richerror.New(op).WithMessage(errmsg.ErrorMsgPhoneNumberIsNotUnique)
	}
	return nil
}
