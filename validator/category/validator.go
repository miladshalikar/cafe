package categoryvalidator

import "context"

type Validator struct {
	repo Repository
}

type Repository interface {
	CheckAreaIsExistByID(ctx context.Context, id uint) (bool, error)
}

func New(r Repository) Validator {
	return Validator{repo: r}
}
