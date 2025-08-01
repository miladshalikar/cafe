package categoryvalidator

import (
	"context"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
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
	const op = "categoryvalidator.checkCategoryIsExist"

	categoryID, ok := value.(uint)
	if !ok {
		return richerror.New(op).WithMessage("invalid category ID").WithKind(richerror.KindInvalid)
	}
	res, err := v.repo.CheckCategoryIsExistByID(ctx, categoryID)
	if err != nil {
		return richerror.New(op).WithWarpError(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	if !res {
		return richerror.New(op).WithMessage("category not found").WithKind(richerror.KindNotFound)
	}
	return nil
}
