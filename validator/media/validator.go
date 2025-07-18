package mediavalidator

import (
	"context"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
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
	mediaID, ok := value.(uint)
	if !ok {
		return validation.NewError("media_id", "media_id must be a valid unsigned integer")
	}
	res, err := v.repo.CheckMediaIsExistByID(ctx, mediaID)
	if err != nil {
		return validation.NewError("media_id", fmt.Sprintf("failed to check media existence: %w", err))
	}
	if !res {
		return validation.NewError("media_id", "media with ID %d not found")
	}
	return nil
}
