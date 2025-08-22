package paymentvalidator

import (
	"context"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

type Validator struct {
	repo Repository
}

type Repository interface {
	CheckPaymentIsExistByID(ctx context.Context, id uint) (bool, error)
}

func New(r Repository) Validator {
	return Validator{repo: r}
}

func (v Validator) checkPaymentIsExist(ctx context.Context, value any) error {
	const op = "paymentvalidator.checkPaymentIsExist"

	itemID, ok := value.(uint)
	if !ok {
		return richerror.New(op).WithMessage("invalid item ID").WithKind(richerror.KindInvalid)
	}
	res, err := v.repo.CheckPaymentIsExistByID(ctx, itemID)
	if err != nil {
		return richerror.New(op).WithWarpError(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	if !res {
		return richerror.New(op).WithMessage("payment not found").WithKind(richerror.KindNotFound)
	}
	return nil
}
