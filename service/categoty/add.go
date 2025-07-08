package categoryservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	categoryparam "github.com/miladshalikar/cafe/param/category"
)

func (s Service) AddNewCategory(ctx context.Context, req categoryparam.AddNewCategoryRequest) (categoryparam.AddNewCategoryResponse, error) {

	cCategory := entity.Category{
		Title: req.Title,
		Logo:  req.Logo,
	}

	category, cErr := s.repo.AddNewCategory(ctx, cCategory)
	if cErr != nil {
		return categoryparam.AddNewCategoryResponse{}, cErr
	}

	return categoryparam.AddNewCategoryResponse{
		Id:    category.Id,
		Title: category.Title,
		Logo:  category.Logo,
	}, nil
}
