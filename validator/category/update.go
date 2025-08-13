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

	if err := validation.ValidateStructWithContext(ctx, &req,
		validation.Field(&req.Title, validation.Required, validation.Length(3, 190)),
		validation.Field(&req.MediaID, validation.When(req.MediaID != 0, validation.WithContext(v.media.CheckMediaIsExistByID))),
	); err != nil {
		fieldErrors := make(map[string]string)
		vErr := validation.Errors{}
		if errors.As(err, &vErr) {
			for key, value := range vErr {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}
		return fieldErrors, richerror.New(op).WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).
			WithWarpError(err)
	}

	return nil, nil
}
