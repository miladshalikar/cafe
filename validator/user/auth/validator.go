package userauthvalidator

import (
	"context"
	"errors"
)

const (
	phoneNumberRegex = `^(?:\+98|0)?9\d{9}$`
	emailRegex       = `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`
)

type Repository interface {
	PhoneNumberExistInDB(ctx context.Context, phoneNumber string) (bool, error)
	EmailExistInDB(ctx context.Context, email string) (bool, error)
}

type Validator struct {
	repo Repository
}

func New(r Repository) Validator {
	return Validator{repo: r}
}

func (v Validator) isEmailExistInDB(value interface{}) error {
	email, ok := value.(string)
	if !ok {
		return errors.New("something went wrong")
	}
	existed, err := v.repo.EmailExistInDB(context.Background(), email)
	if err != nil {
		return errors.New("something went wrong")
	}
	if existed {
		return errors.New("user exist")
	}
	return nil
}

func (v Validator) isPhoneNumberExistInDB(value interface{}) error {
	phoneNumber, ok := value.(string)
	if !ok {
		return errors.New("something went wrong")
	}
	existed, err := v.repo.PhoneNumberExistInDB(context.Background(), phoneNumber)
	if err != nil {
		return errors.New("something went wrong")
	}
	if existed {
		return errors.New("user exist")
	}
	return nil
}
