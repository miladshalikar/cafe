package categoryservice

import (
	"context"
	"errors"
	categoryparam "github.com/miladshalikar/cafe/param/category"
	mediaparam "github.com/miladshalikar/cafe/param/media"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (s Service) ShowCategory(ctx context.Context, req categoryparam.ShowSingleCategoryRequest) (categoryparam.ShowSingleCategoryResponse, error) {
	const op = "categoryservice.ShowCategory"

	category, cErr := s.repo.GetCategoryByID(ctx, req.ID)
	if cErr != nil {
		return categoryparam.ShowSingleCategoryResponse{}, richerror.New(op).WithWarpError(cErr)
	}

	var url string
	if category.MediaID != 0 {
		media, mErr := s.client.GetURLMedia(ctx, mediaparam.GetURLRequest{ID: category.MediaID})
		if mErr != nil {
			var rErr richerror.RichError
			if errors.As(mErr, &rErr) && rErr.Kind() == richerror.KindNotFound {
				url = ""
			}
			return categoryparam.ShowSingleCategoryResponse{}, richerror.New(op).WithWarpError(mErr)
		}
		url = media.URL
	}

	return categoryparam.ShowSingleCategoryResponse{
		CategoryInfo: categoryparam.CategoryInfo{
			ID:      category.ID,
			Title:   category.Title,
			MediaID: category.MediaID,
			URL:     url,
		},
	}, nil
}
