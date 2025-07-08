package categoryservice

import (
	"context"
	categoryparam "github.com/miladshalikar/cafe/param/category"
)

func (s Service) DeleteCategory(ctx context.Context, req categoryparam.DeleteCategoryRequest) (categoryparam.DeleteCategoryResponse, error) {

	if cErr := s.repo.DeleteCategory(ctx, req.ID); cErr != nil {
		return categoryparam.DeleteCategoryResponse{}, cErr
	}

	return categoryparam.DeleteCategoryResponse{}, nil
}
