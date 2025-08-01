package itemvalidator

import (
	"context"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	itemparam "github.com/miladshalikar/cafe/param/item"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (v Validator) ValidateDeleteItem(ctx context.Context, req itemparam.DeleteItemRequest) (map[string]string, error) {
	const op = "itemvalidator.ValidateDeleteItem"

	if err := validation.ValidateStructWithContext(ctx, &req,
		validation.Field(&req.ID, validation.Required, validation.WithContext(v.checkItemIsExist)),
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
