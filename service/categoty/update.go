package categoryservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	categoryparam "github.com/miladshalikar/cafe/param/category"
	mediaparam "github.com/miladshalikar/cafe/param/media"
)

func (s Service) UpdateCategory(ctx context.Context, req categoryparam.UpdateCategoryRequest) (categoryparam.UpdateCategoryResponse, error) {

	currentCategory, cErr := s.repo.GetCategoryByID(ctx, req.ID)
	if cErr != nil {
		return categoryparam.UpdateCategoryResponse{}, cErr
	}

	if currentCategory.MediaID != req.MediaID {

		if _, sErr := s.Client.DeleteMedia(ctx, mediaparam.DeleteMediaRequest{ID: currentCategory.MediaID}); sErr != nil {
			return categoryparam.UpdateCategoryResponse{}, sErr
		}
	}

	category := entity.Category{
		ID:      req.ID,
		Title:   req.Title,
		MediaID: req.MediaID,
	}

	if rErr := s.repo.UpdateCategory(ctx, category); rErr != nil {
		return categoryparam.UpdateCategoryResponse{}, rErr
	}
	mediaURL, uErr := s.Client.GetURLMedia(ctx, mediaparam.GetURLRequest{ID: category.MediaID})
	if uErr != nil {
		return categoryparam.UpdateCategoryResponse{}, uErr
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
