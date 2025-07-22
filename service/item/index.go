package itemservice

import (
	"context"
	commonparam "github.com/miladshalikar/cafe/param/common"
	itemparam "github.com/miladshalikar/cafe/param/item"
	mediaparam "github.com/miladshalikar/cafe/param/media"
)

func (s Service) GetItems(ctx context.Context, req itemparam.GetItemRequest) (itemparam.GetItemResponse, error) {

	total, tErr := s.repo.GetTotalCountItem(ctx, req.Search.Search)
	if tErr != nil {
		return itemparam.GetItemResponse{}, tErr
	}

	items, cErr := s.repo.GetItemsWithPagination(ctx, req.Pagination.GetPageSize(), req.Pagination.GetOffset(), req.Search.Search)
	if cErr != nil {
		return itemparam.GetItemResponse{}, cErr
	}

	var mediaIDs []uint
	for _, item := range items {
		mediaIDs = append(mediaIDs, item.MediaID)
	}

	cachedURLs, err := s.cache.MGetMediaURLs(ctx, mediaIDs)
	if err != nil {
		return itemparam.GetItemResponse{}, err
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
			return itemparam.GetItemResponse{}, mErr
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

	return itemparam.GetItemResponse{
		Pagination: commonparam.PaginationResponse{
			PageSize:   req.Pagination.GetPageSize(),
			PageNumber: req.Pagination.GetPageNumber(),
			Total:      total,
		},
		Items: itemsInfo,
	}, nil

}
