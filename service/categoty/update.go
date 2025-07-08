package categoryservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	categoryparam "github.com/miladshalikar/cafe/param/category"
)

func (s Service) UpdateCategory(ctx context.Context, req categoryparam.UpdateCategoryRequest) (categoryparam.UpdateCategoryResponse, error) {
	category := entity.Category{
		Id:    req.ID,
		Title: req.Title,
		Logo:  req.Logo,
	}

	if cErr := s.repo.UpdateCategory(ctx, category); cErr != nil {
		return categoryparam.UpdateCategoryResponse{}, cErr
	}

	return categoryparam.UpdateCategoryResponse{
		ID:    category.Id,
		Title: category.Title,
		Logo:  category.Logo,
	}, nil
}
