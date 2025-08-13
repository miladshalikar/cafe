package itemservice

import (
	"context"
	"errors"
	commonparam "github.com/miladshalikar/cafe/param/common"
	itemparam "github.com/miladshalikar/cafe/param/item"
	mediaparam "github.com/miladshalikar/cafe/param/media"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (s Service) GetItems(ctx context.Context, req itemparam.GetItemsRequest) (itemparam.GetItemsResponse, error) {
	const op = "itemservice.GetItems"

	total, tErr := s.repo.GetTotalCountItemWithSearchAndFilter(ctx, req.Search, req.Filter)
	if tErr != nil {
		return itemparam.GetItemsResponse{}, richerror.New(op).WithWarpError(tErr)
	}

	items, cErr := s.repo.GetItemsWithPaginationAndSearchAndFilter(ctx, req.Pagination, req.Search, req.Filter)
	if cErr != nil {
		return itemparam.GetItemsResponse{}, richerror.New(op).WithWarpError(cErr)
	}

	var itemsInfo []itemparam.ItemInfo
	if len(items) > 0 {
		var mediaIDs []uint
		for _, item := range items {
			mediaIDs = append(mediaIDs, item.MediaID)
		}

		cachedURLs, err := s.cache.MGetMediaURLs(ctx, mediaIDs)
		if err != nil {
			return itemparam.GetItemsResponse{}, richerror.New(op).WithWarpError(err)
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

				return itemparam.GetItemsResponse{}, richerror.New(op).WithWarpError(mErr)
			}

			sErr := s.cache.SetMediaURLByMediaID(ctx, id, mediaRes.URL)
			if sErr != nil {
				//به جای برگرداندن خطا لاگ کنیم؟todo
				return itemparam.GetItemsResponse{}, richerror.New(op).WithWarpError(sErr)
			}

			cachedURLs[id] = mediaRes.URL

		}

		for _, item := range items {
			itemsInfo = append(itemsInfo, itemparam.ItemInfo{
				ID:          item.ID,
				Title:       item.Title,
				Description: item.Description,
				Price:       item.Price,
				Quantity:    item.Quantity,
				CategoryID:  item.CategoryID,
				MediaID:     item.MediaID,
				URL:         cachedURLs[item.MediaID],
			})
		}
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
