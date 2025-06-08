package userauthvalidator

import (
	"context"
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
