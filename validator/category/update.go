package categoryvalidator

import (
	"context"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	categoryparam "github.com/miladshalikar/cafe/param/category"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (v Validator) ValidateUpdateCategory(ctx context.Context, req categoryparam.UpdateCategoryRequest) (map[string]string, error) {
	const op = "categoryvalidator.ValidateUpdateCategory"

	fieldErrors := make(map[string]string)

	if err := validation.ValidateStructWithContext(ctx, &req,
		validation.Field(&req.Title, validation.Required, validation.Length(3, 190)),
		validation.Field(&req.MediaID, validation.When(req.MediaID != 0, validation.WithContext(v.media.CheckMediaIsExistByID))),
	); err != nil {
		vErr := validation.Errors{}
		if errors.As(err, &vErr) {
			for key, value := range vErr {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}
	}

	currentCategory, err := v.repo.GetCategoryByID(ctx, req.ID)
	if err != nil {
		return nil, richerror.New(op).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).
			WithKind(richerror.KindUnexpected).
			WithMeta(map[string]interface{}{"req": req}).
			WithWarpError(err)
	}

	if req.Title != currentCategory.Title {
		exists, cErr := v.repo.CheckCategoryIsExistByTitle(ctx, req.Title)
		if cErr != nil {
			return nil, richerror.New(op).
				WithWarpError(cErr).
				WithMessage(errmsg.ErrorMsgSomethingWentWrong).
				WithKind(richerror.KindUnexpected).
				WithMeta(map[string]interface{}{"req": req})
		}
		if exists {
			fieldErrors["title"] = "category name already exists"
		}
	}

	if len(fieldErrors) == 0 {
		return nil, nil
	}

	return fieldErrors, richerror.New(op).WithMessage(errmsg.ErrorMsgInvalidInput).
		WithKind(richerror.KindInvalid).
		WithMeta(map[string]interface{}{"req": req}).
		WithWarpError(err)

}
