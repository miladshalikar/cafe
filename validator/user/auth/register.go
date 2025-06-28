package userauthvalidator

import (
	"context"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	param "github.com/miladshalikar/cafe/param/user/authservice"
	"regexp"
)

func (v Validator) ValidateRegisterRequest(ctx context.Context, req param.RegisterRequest) (map[string]string, error) {

	if err := validation.ValidateStructWithContext(ctx, &req,
		validation.Field(&req.FirstName, validation.Required, validation.Length(2, 50)),
		validation.Field(&req.LastName, validation.Required, validation.Length(2, 50)),
		validation.Field(&req.Email,
			validation.Required,
			validation.Length(5, 100),
			validation.Match(regexp.MustCompile(emailRegex)).Error("invalid email address"),
			validation.WithContext(v.isEmailExistInDB)),
		validation.Field(&req.PhoneNumber,
			validation.Required,
			validation.Length(5, 100),
			validation.Match(regexp.MustCompile(phoneNumberRegex)).Error("invalid phone number"),
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
		return fieldErrors, err
	}
	return nil, nil
}

func (v Validator) isEmailExistInDB(ctx context.Context, value interface{}) error {
	email, ok := value.(string)
	if !ok {
		return errors.New("something went wrong1")
	}
	existed, err := v.repo.EmailExistInDB(ctx, email)
	if err != nil {
		return errors.New("something went wrong2")
	}
	if existed {
		return errors.New("email exist")
	}
	return nil
}

func (v Validator) isPhoneNumberExistInDB(ctx context.Context, value interface{}) error {
	phoneNumber, ok := value.(string)
	if !ok {
		return errors.New("something went wrong")
	}
	existed, err := v.repo.PhoneNumberExistInDB(ctx, phoneNumber)
	if err != nil {
		return errors.New("something went wrong")
	}
	if existed {
		return errors.New("phoneNumber exist")
	}
	return nil
}
