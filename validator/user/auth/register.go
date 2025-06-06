package userauthvalidator

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	param "github.com/miladshalikar/cafe/param/authservice"
	"regexp"
)

func (v Validator) ValidateRegisterRequest(req param.RegisterRequest) (map[string]string, error) {

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.FirstName, validation.Required, validation.Length(2, 50)),
		validation.Field(&req.LastName, validation.Required, validation.Length(2, 50)),
		validation.Field(&req.Email,
			validation.Required,
			validation.Length(5, 100),
			validation.Match(regexp.MustCompile(emailRegex)).Error("invalid email address"),
			validation.By(v.isEmailExistInDB)),
		validation.Field(&req.PhoneNumber,
			validation.Required,
			validation.Length(5, 100),
			validation.Match(regexp.MustCompile(phoneNumberRegex)).Error("invalid phone number"),
			validation.By(v.isPhoneNumberExistInDB)),
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
