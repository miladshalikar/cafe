package categoryvalidator

import (
	"context"
	"errors"
)

type Validator struct {
	repo Repository
}

type Repository interface {
	CheckCategoryIsExistByID(ctx context.Context, id uint) (bool, error)
}

func New(r Repository) Validator {
	return Validator{repo: r}
}

func (v Validator) checkCategoryIsExist(ctx context.Context, value any) error {
	categoryID, ok := value.(uint)
	if !ok {
		return errors.New("categoryID should be uint")
	}
	res, err := v.repo.CheckCategoryIsExistByID(ctx, categoryID)
	if err != nil {
		return err
	}
	if !res {
		return errors.New("category is not exist")
	}
	return nil
}
