package itemservice

import (
	"context"
	commonparam "github.com/miladshalikar/cafe/param/common"
	itemparam "github.com/miladshalikar/cafe/param/item"
	mediaparam "github.com/miladshalikar/cafe/param/media"
)

func (s Service) GetItems(ctx context.Context, req itemparam.GetItemsRequest) (itemparam.GetItemsResponse, error) {

	total, tErr := s.repo.GetTotalCountItemWithSearchAndFilter(ctx, req.Search, req.Filter)
	if tErr != nil {
		return itemparam.GetItemsResponse{}, tErr
	}

	items, cErr := s.repo.GetItemsWithPaginationAndSearchAndFilter(ctx, req.Pagination, req.Search, req.Filter)
	if cErr != nil {
		return itemparam.GetItemsResponse{}, cErr
	}

	var mediaIDs []uint
	for _, item := range items {
		mediaIDs = append(mediaIDs, item.MediaID)
	}

	cachedURLs, err := s.cache.MGetMediaURLs(ctx, mediaIDs)
	if err != nil {
		return itemparam.GetItemsResponse{}, err
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
			return itemparam.GetItemsResponse{}, mErr
		}

		_ = s.cache.SetMediaURLByMediaID(ctx, id, mediaRes.URL)

		cachedURLs[id] = mediaRes.URL

	}

	var itemsInfo []itemparam.ItemInfo
	for _, item := range items {
		itemsInfo = append(itemsInfo, itemparam.ItemInfo{
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			Price:       item.Price,
			CategoryID:  item.CategoryID,
			MediaID:     item.MediaID,
			URL:         cachedURLs[item.MediaID],
		})
	}

	return itemparam.GetItemsResponse{
		Pagination: commonparam.PaginationResponse{
			PageSize:   req.Pagination.GetPageSize(),
			PageNumber: req.Pagination.GetPageNumber(),
			Total:      total,
		},
		Items: itemsInfo,
	}, nil

}
