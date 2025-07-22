package categoryservice

import (
	"context"
	categoryparam "github.com/miladshalikar/cafe/param/category"
	mediaparam "github.com/miladshalikar/cafe/param/media"
)

func (s Service) DeleteCategory(ctx context.Context, req categoryparam.DeleteCategoryRequest) (categoryparam.DeleteCategoryResponse, error) {

	category, mErr := s.repo.GetCategoryByID(ctx, req.ID)
	if mErr != nil {
		return categoryparam.DeleteCategoryResponse{}, mErr
	}

	if cErr := s.repo.DeleteCategory(ctx, req.ID); cErr != nil {
		return categoryparam.DeleteCategoryResponse{}, cErr
	}

	if _, dErr := s.Client.DeleteMedia(ctx, mediaparam.DeleteMediaRequest{ID: category.MediaID}); dErr != nil {

		_ = s.repo.UndoDeleteCategory(ctx, req.ID)

		return categoryparam.DeleteCategoryResponse{}, dErr
	}

	return categoryparam.DeleteCategoryResponse{}, nil
}
