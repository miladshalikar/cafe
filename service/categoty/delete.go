package categoryservice

import (
	"context"
	categoryparam "github.com/miladshalikar/cafe/param/category"
	mediaparam "github.com/miladshalikar/cafe/param/media"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (s Service) DeleteCategory(ctx context.Context, req categoryparam.DeleteCategoryRequest) (categoryparam.DeleteCategoryResponse, error) {
	const op = "categoryservice.DeleteCategory"

	category, mErr := s.repo.GetCategoryByID(ctx, req.ID)
	if mErr != nil {
		return categoryparam.DeleteCategoryResponse{}, richerror.New(op).WithWarpError(mErr)
	}

	//todo

	if cErr := s.repo.DeleteCategory(ctx, req.ID); cErr != nil {
		return categoryparam.DeleteCategoryResponse{}, richerror.New(op).WithWarpError(cErr)
	}

	if _, dErr := s.client.DeleteMedia(ctx, mediaparam.DeleteMediaRequest{ID: category.MediaID}); dErr != nil {

		_ = s.repo.UndoDeleteCategory(ctx, req.ID)

		return categoryparam.DeleteCategoryResponse{}, richerror.New(op).WithWarpError(dErr).WithMeta(map[string]interface{}{"media_id": category.MediaID})
	}

	return categoryparam.DeleteCategoryResponse{}, nil
}
