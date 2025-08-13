package categoryservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	categoryparam "github.com/miladshalikar/cafe/param/category"
	mediaparam "github.com/miladshalikar/cafe/param/media"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (s Service) UpdateCategory(ctx context.Context, req categoryparam.UpdateCategoryRequest) (categoryparam.UpdateCategoryResponse, error) {
	const op = "categoryservice.UpdateCategory"

	currentCategory, cErr := s.repo.GetCategoryByID(ctx, req.ID)
	if cErr != nil {
		return categoryparam.UpdateCategoryResponse{}, richerror.New(op).WithWarpError(cErr)
	}

	if currentCategory.MediaID != req.MediaID && req.MediaID != 0 {

		if _, sErr := s.client.DeleteMedia(ctx, mediaparam.DeleteMediaRequest{ID: currentCategory.MediaID}); sErr != nil {
			return categoryparam.UpdateCategoryResponse{}, richerror.New(op).WithWarpError(sErr)
		}
	}

	category := entity.Category{
		ID:    req.ID,
		Title: req.Title,
	}
	if req.MediaID == 0 {
		category.MediaID = currentCategory.MediaID
	} else {
		category.MediaID = req.MediaID
	}

	if rErr := s.repo.UpdateCategory(ctx, category); rErr != nil {
		return categoryparam.UpdateCategoryResponse{}, richerror.New(op).WithWarpError(rErr)
	}
	mediaURL, uErr := s.client.GetURLMedia(ctx, mediaparam.GetURLRequest{ID: category.MediaID})
	if uErr != nil {
		return categoryparam.UpdateCategoryResponse{}, richerror.New(op).WithWarpError(uErr)
	}

	return categoryparam.UpdateCategoryResponse{
		Category: categoryparam.CategoryInfo{
			ID:      category.ID,
			Title:   category.Title,
			MediaID: category.MediaID,
			URL:     mediaURL.URL,
		},
	}, nil
}
