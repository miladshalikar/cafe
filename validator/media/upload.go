package mediavalidator

import (
	"context"
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/miladshalikar/cafe/entity"
	mediaparam "github.com/miladshalikar/cafe/param/media"
	"github.com/miladshalikar/cafe/pkg/list"
	"path/filepath"
	"strings"
)

func (v Validator) ValidateUploadFile(ctx context.Context, req mediaparam.UploadMediaRequest) (map[string]string, error) {

	if err := validation.ValidateStructWithContext(ctx, &req,
		validation.Field(&req.Bucket, validation.Required),
		validation.Field(&req.Size, validation.Max(entity.MaxFileUploadSize)),
		validation.Field(&req.Filename, validation.By(ValidateFileExtension)),
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

func ValidateFileExtension(value interface{}) error {
	filename, ok := value.(string)
	if !ok {
		return validation.NewError("file_extension", "value is not a string")
	}

	ext := strings.ToLower(filepath.Ext(filename)[1:])

	if !list.CheckStringInList(ext, entity.ValidExt) {
		return validation.NewError("file_extension", fmt.Sprintf("valid extensions are: %v", entity.ValidExt))
	}

	return nil
}
