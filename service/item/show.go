package itemservice

import (
	"context"
	itemparam "github.com/miladshalikar/cafe/param/item"
	mediaparam "github.com/miladshalikar/cafe/param/media"
)

func (s Service) ShowItem(ctx context.Context, req itemparam.ShowSingleItemRequest) (itemparam.ShowSingleItemResponse, error) {

	item, cErr := s.repo.GetItemByID(ctx, req.ID)
	if cErr != nil {
		return itemparam.ShowSingleItemResponse{}, cErr
	}

	media, mErr := s.client.GetURLMedia(ctx, mediaparam.GetURLRequest{ID: item.MediaID})
	if mErr != nil {
		return itemparam.ShowSingleItemResponse{}, mErr
	}

	return itemparam.ShowSingleItemResponse{
		ItemInfo: itemparam.ItemInfo{
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			Price:       item.Price,
			CategoryID:  item.CategoryID,
			MediaID:     item.MediaID,
			URL:         media.URL,
		},
	}, nil
}
