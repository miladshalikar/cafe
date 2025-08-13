package itemvalidator

import (
	"context"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

type Validator struct {
	repo  Repository
	media Media
}

type Repository interface {
	CheckItemIsExistByID(ctx context.Context, id uint) (bool, error)
}

type Media interface {
	CheckMediaIsExistByID(ctx context.Context, value any) error
}

func New(r Repository, m Media) Validator {
	return Validator{repo: r, media: m}
}

func (v Validator) checkItemIsExist(ctx context.Context, value any) error {
	const op = "itemvalidator.checkItemIsExist"

	itemID, ok := value.(uint)
	if !ok {
		return richerror.New(op).WithMessage("invalid item ID").WithKind(richerror.KindInvalid)
	}
	res, err := v.repo.CheckItemIsExistByID(ctx, itemID)
	if err != nil {
		return richerror.New(op).WithWarpError(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	if !res {
		return richerror.New(op).WithMessage("item not found").WithKind(richerror.KindNotFound)
	}
	return nil
}
