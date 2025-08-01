package mediavalidator

import (
	"context"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/miladshalikar/cafe/entity"
	mediaparam "github.com/miladshalikar/cafe/param/media"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/list"
	"github.com/miladshalikar/cafe/pkg/richerror"
	"path/filepath"
	"strings"
)

func (v Validator) ValidateUploadFile(ctx context.Context, req mediaparam.UploadMediaRequest) (map[string]string, error) {
	const op = "mediavalidator.ValidateUploadFile"

	if err := validation.ValidateStructWithContext(ctx, &req,
		validation.Field(&req.Size, validation.Max(entity.MaxFileUploadSize).Error(errmsg.ErrorFileSize)),
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
		return fieldErrors, richerror.New(op).WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).
			WithWarpError(err)
	}

	return nil, nil
}

func ValidateFileExtension(value interface{}) error {
	const op = "mediavalidator.ValidateFileExtension"

	filename, ok := value.(string)
	if !ok {
		return richerror.New(op).WithMessage("invalid file format").WithKind(richerror.KindInvalid)
	}

	ext := strings.ToLower(filepath.Ext(filename)[1:])

	if !list.CheckStringInList(ext, entity.ValidExt) {
		return richerror.New(op).WithMessage("invalid file format").WithKind(richerror.KindInvalid)
	}

	return nil
}
