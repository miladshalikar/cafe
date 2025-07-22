package itemvalidator

import (
	"context"
	"errors"
)

type Validator struct {
	repo Repository
}

type Repository interface {
	CheckItemIsExistByID(ctx context.Context, id uint) (bool, error)
}

func New(r Repository) Validator {
	return Validator{repo: r}
}

func (v Validator) checkItemIsExist(ctx context.Context, value any) error {
	itemID, ok := value.(uint)
	if !ok {
		return errors.New("itemID should be uint")
	}
	res, err := v.repo.CheckItemIsExistByID(ctx, itemID)
	if err != nil {
		return err
	}
	if !res {
		return errors.New("item is not exist")
	}
	return nil
}
