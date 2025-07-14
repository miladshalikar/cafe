package categoryservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	categoryparam "github.com/miladshalikar/cafe/param/category"
)

func (s Service) UpdateCategory(ctx context.Context, req categoryparam.UpdateCategoryRequest) (categoryparam.UpdateCategoryResponse, error) {
	category := entity.Category{
		ID:    req.ID,
		Title: req.Title,
	}

	if cErr := s.repo.UpdateCategory(ctx, category); cErr != nil {
		return categoryparam.UpdateCategoryResponse{}, cErr
	}

	return categoryparam.UpdateCategoryResponse{
		Category: categoryparam.CategoryInfo{
			ID:    category.ID,
			Title: category.Title,
		},
	}, nil
}
