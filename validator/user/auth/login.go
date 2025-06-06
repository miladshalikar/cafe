package userauthvalidator

import (
	"context"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	param "github.com/miladshalikar/cafe/param/authservice"
	"regexp"
)

func (v Validator) ValidateLoginWithEmailRequest(req param.LoginWithEmailRequest) (map[string]string, error) {
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Password, validation.Required, validation.NotNil),
		validation.Field(&req.Email,
			validation.Required,
			validation.Match(regexp.MustCompile(emailRegex)).Error("invalid email address"),
			validation.By(v.isUserExist))); err != nil {

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

func (v Validator) isUserExist(value interface{}) error {

	email, ok := value.(string)
	if !ok {
		return errors.New("something went wrong")
	}

	existed, err := v.repo.EmailExistInDB(context.Background(), email)
	if err != nil {
		return errors.New("user not exist")
	}

	if !existed {
		return errors.New("user not exist")
	}
	return nil
}
