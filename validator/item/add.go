package itemvalidator

import (
	"context"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	itemparam "github.com/miladshalikar/cafe/param/item"
)

func (v Validator) ValidateAddItem(ctx context.Context, req itemparam.AddNewItemRequest) (map[string]string, error) {
	if err := validation.ValidateStructWithContext(ctx, &req,
		validation.Field(&req.Title, validation.Required, validation.Length(3, 190)),
		validation.Field(&req.Description, validation.NilOrNotEmpty),
		validation.Field(&req.Price, validation.Required, validation.Min(0.0)),
		validation.Field(&req.CategoryID, validation.Required, validation.Min(1)),
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
