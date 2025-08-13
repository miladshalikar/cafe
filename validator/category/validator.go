package categoryvalidator

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

type Validator struct {
	repo  Repository
	media Media
}

type Repository interface {
	CheckCategoryIsExistByID(ctx context.Context, id uint) (bool, error)
	CheckCategoryIsExistByTitle(ctx context.Context, title string) (bool, error)
	GetCategoryByID(ctx context.Context, id uint) (entity.Category, error)
}

type Media interface {
	CheckMediaIsExistByID(ctx context.Context, value any) error
}

func New(r Repository, m Media) Validator {
	return Validator{repo: r, media: m}
}

func (v Validator) checkCategoryIsExistByID(ctx context.Context, value any) error {
	const op = "categoryvalidator.checkCategoryIsExistByID"

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

func (v Validator) checkCategoryIsExistByTitle(ctx context.Context, value any) error {
	const op = "categoryvalidator.checkCategoryIsExistByTitle"

	categoryTitle, ok := value.(string)
	if !ok {
		return richerror.New(op).WithMessage("invalid category Title").WithKind(richerror.KindInvalid)
	}
	res, err := v.repo.CheckCategoryIsExistByTitle(ctx, categoryTitle)
	if err != nil {
		return richerror.New(op).WithWarpError(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	if res {
		return richerror.New(op).WithMessage("category name already exists").WithKind(richerror.KindNotFound)
	}
	return nil
}
