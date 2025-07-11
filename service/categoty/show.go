package categoryservice

import (
	"context"
	categoryparam "github.com/miladshalikar/cafe/param/category"
)

func (s Service) ShowCategory(ctx context.Context, req categoryparam.ShowSingleCategoryRequest) (categoryparam.ShowSingleCategoryResponse, error) {

	category, cErr := s.repo.GetCategoryById(ctx, req.ID)
	if cErr != nil {
		return categoryparam.ShowSingleCategoryResponse{}, cErr
	}

	return categoryparam.ShowSingleCategoryResponse{
		ID:    category.ID,
		Title: category.Title,
		Logo:  category.Logo,
	}, nil
}
