package categoryservice

import (
	"context"
	categoryparam "github.com/miladshalikar/cafe/param/category"
	mediaparam "github.com/miladshalikar/cafe/param/media"
)

func (s Service) ShowCategory(ctx context.Context, req categoryparam.ShowSingleCategoryRequest) (categoryparam.ShowSingleCategoryResponse, error) {

	category, cErr := s.repo.GetCategoryByID(ctx, req.ID)
	if cErr != nil {
		return categoryparam.ShowSingleCategoryResponse{}, cErr
	}

	media, mErr := s.Client.GetURLMedia(ctx, mediaparam.GetURLRequest{ID: category.MediaID})
	if mErr != nil {
		return categoryparam.ShowSingleCategoryResponse{}, mErr
	}

	return categoryparam.ShowSingleCategoryResponse{
		CategoryInfo: categoryparam.CategoryInfo{
			ID:      category.ID,
			Title:   category.Title,
			MediaID: category.MediaID,
			URL:     media.URL,
		},
	}, nil
}
