package categoryservice

import (
	"context"
	"errors"
	categoryparam "github.com/miladshalikar/cafe/param/category"
	commonparam "github.com/miladshalikar/cafe/param/common"
	mediaparam "github.com/miladshalikar/cafe/param/media"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (s Service) GetCategories(ctx context.Context, req categoryparam.GetCategoriesRequest) (categoryparam.GetCategoriesResponse, error) {
	const op = "categoryservice.GetCategories"

	total, tErr := s.repo.GetTotalCountCategoryWithSearch(ctx, req.Search)
	if tErr != nil {
		return categoryparam.GetCategoriesResponse{}, richerror.New(op).WithWarpError(tErr)
	}

	categories, cErr := s.repo.GetCategoriesWithPaginationAndSearch(ctx, req.Pagination, req.Search)
	if cErr != nil {
		return categoryparam.GetCategoriesResponse{}, richerror.New(op).WithWarpError(cErr)
	}

	var categoriesInfo []categoryparam.CategoryInfo
	if len(categories) > 0 {
		var mediaIDs []uint
		for _, category := range categories {
			mediaIDs = append(mediaIDs, category.MediaID)
		}

		cachedURLs, err := s.cache.MGetMediaURLs(ctx, mediaIDs)
		if err != nil {
			return categoryparam.GetCategoriesResponse{}, richerror.New(op).WithWarpError(err)
		}

		missingIDs := make([]uint, 0)
		for _, id := range mediaIDs {
			if _, found := cachedURLs[id]; !found {
				missingIDs = append(missingIDs, id)
			}
		}

		for _, id := range missingIDs {

			if id == 0 {
				cachedURLs[id] = ""
				continue
			}

			mediaRes, mErr := s.client.GetURLMedia(ctx, mediaparam.GetURLRequest{ID: id})
			if mErr != nil {

				var rErr richerror.RichError
				if errors.As(mErr, &rErr) && rErr.Kind() == richerror.KindNotFound {
					cachedURLs[id] = ""
					continue
				}

				return categoryparam.GetCategoriesResponse{}, richerror.New(op).WithWarpError(mErr)
			}

			sErr := s.cache.SetMediaURLByMediaID(ctx, id, mediaRes.URL)
			if sErr != nil {
				//به جای برگرداندن خطا لاگ کنیم؟todo
				return categoryparam.GetCategoriesResponse{}, richerror.New(op).WithWarpError(sErr)
			}

			cachedURLs[id] = mediaRes.URL
		}

		for _, category := range categories {
			categoriesInfo = append(categoriesInfo, categoryparam.CategoryInfo{
				ID:      category.ID,
				Title:   category.Title,
				MediaID: category.MediaID,
				URL:     cachedURLs[category.MediaID],
			})
		}
	}

	return categoryparam.GetCategoriesResponse{
		Pagination: commonparam.PaginationResponse{
			PageSize:   req.Pagination.GetPageSize(),
			PageNumber: req.Pagination.GetPageNumber(),
			Total:      total,
		},
		Categories: categoriesInfo,
	}, nil

}
