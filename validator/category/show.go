package categoryvalidator

import (
	"context"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	categoryparam "github.com/miladshalikar/cafe/param/category"
)

func (v Validator) ValidateShowSingleCategory(ctx context.Context, req categoryparam.ShowSingleCategoryRequest) (map[string]string, error) {

	if err := validation.ValidateStructWithContext(ctx, &req,
		validation.Field(&req.ID, validation.Required, validation.WithContext(v.checkCategoryIsExist)),
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
