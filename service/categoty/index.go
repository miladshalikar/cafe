package categoryservice

import (
	"context"
	categoryparam "github.com/miladshalikar/cafe/param/category"
	commonparam "github.com/miladshalikar/cafe/param/common"
	mediaparam "github.com/miladshalikar/cafe/param/media"
)

func (s Service) GetCategories(ctx context.Context, req categoryparam.GetCategoriesRequest) (categoryparam.GetCategoriesResponse, error) {

	total, tErr := s.repo.GetTotalCountCategoryWithSearch(ctx, req.Search)
	if tErr != nil {
		return categoryparam.GetCategoriesResponse{}, tErr
	}

	categories, cErr := s.repo.GetCategoriesWithPaginationAndSearch(ctx, req.Pagination, req.Search)
	if cErr != nil {
		return categoryparam.GetCategoriesResponse{}, cErr
	}

	var mediaIDs []uint
	for _, category := range categories {
		mediaIDs = append(mediaIDs, category.MediaID)
	}

	cachedURLs, err := s.cache.MGetMediaURLs(ctx, mediaIDs)
	if err != nil {
		return categoryparam.GetCategoriesResponse{}, err
	}

	missingIDs := make([]uint, 0)
	for _, id := range mediaIDs {
		if _, found := cachedURLs[id]; !found {
			missingIDs = append(missingIDs, id)
		}
	}

	for _, id := range missingIDs {

		mediaRes, mErr := s.client.GetURLMedia(ctx, mediaparam.GetURLRequest{ID: id})
		if mErr != nil {
			return categoryparam.GetCategoriesResponse{}, mErr
		}

		_ = s.cache.SetMediaURLByMediaID(ctx, id, mediaRes.URL)

		cachedURLs[id] = mediaRes.URL

	}

	var categoriesInfo []categoryparam.CategoryInfo
	for _, category := range categories {
		categoriesInfo = append(categoriesInfo, categoryparam.CategoryInfo{
			ID:      category.ID,
			Title:   category.Title,
			MediaID: category.MediaID,
			URL:     cachedURLs[category.MediaID],
		})
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
