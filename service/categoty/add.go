package categoryservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	categoryparam "github.com/miladshalikar/cafe/param/category"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (s Service) AddNewCategory(ctx context.Context, req categoryparam.AddNewCategoryRequest) (categoryparam.AddNewCategoryResponse, error) {
	const op = "categoryservice.AddNewCategory"

	cCategory := entity.Category{
		Title:   req.Title,
		MediaID: req.MediaID,
	}

	category, cErr := s.repo.AddNewCategory(ctx, cCategory)
	if cErr != nil {
		return categoryparam.AddNewCategoryResponse{}, richerror.New(op).WithWarpError(cErr)
	}

	return categoryparam.AddNewCategoryResponse{
		ID:      category.ID,
		Title:   category.Title,
		MediaID: category.MediaID,
	}, nil
}
