package mediavalidator

import (
	"context"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

type Validator struct {
	repo Repository
}

type Repository interface {
	CheckMediaIsExistByID(ctx context.Context, id uint) (bool, error)
}

func New(r Repository) Validator {
	return Validator{repo: r}
}

func (v Validator) checkMediaIsExistByID(ctx context.Context, value any) error {
	const op = "mediavalidator.checkMediaIsExistByID"

	mediaID, ok := value.(uint)
	if !ok {
		return richerror.New(op).WithMessage("invalid media ID").WithKind(richerror.KindInvalid)
	}
	res, err := v.repo.CheckMediaIsExistByID(ctx, mediaID)
	if err != nil {
		return richerror.New(op).WithWarpError(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	if !res {
		return richerror.New(op).WithMessage("media not found").WithKind(richerror.KindNotFound)
	}
	return nil
}
