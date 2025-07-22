package categoryvalidator

import (
	"context"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	categoryparam "github.com/miladshalikar/cafe/param/category"
)

func (v Validator) ValidateUpdateCategory(ctx context.Context, req categoryparam.UpdateCategoryRequest) (map[string]string, error) {

	if err := validation.ValidateStructWithContext(ctx, &req,
		validation.Field(&req.Title, validation.Required, validation.Length(3, 190)),
		validation.Field(&req.MediaID, validation.Min(0)),
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
		return fieldErrors, err
	}

	return nil, nil
}
