package itemservice

import (
	"context"
	"errors"
	itemparam "github.com/miladshalikar/cafe/param/item"
	mediaparam "github.com/miladshalikar/cafe/param/media"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (s Service) ShowItem(ctx context.Context, req itemparam.ShowSingleItemRequest) (itemparam.ShowSingleItemResponse, error) {
	const op = "itemservice.ShowItem"

	item, cErr := s.repo.GetItemByID(ctx, req.ID)
	if cErr != nil {
		return itemparam.ShowSingleItemResponse{}, richerror.New(op).WithWarpError(cErr)
	}

	var url string
	if item.MediaID != 0 {
		media, mErr := s.client.GetURLMedia(ctx, mediaparam.GetURLRequest{ID: item.MediaID})
		if mErr != nil {
			var rErr richerror.RichError
			if errors.As(mErr, &rErr) && rErr.Kind() == richerror.KindNotFound {
				url = ""
			}
			return itemparam.ShowSingleItemResponse{}, richerror.New(op).WithWarpError(mErr)
		}
		url = media.URL
	}

	return itemparam.ShowSingleItemResponse{
		ItemInfo: itemparam.ItemInfo{
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			Price:       item.Price,
			Quantity:    item.Quantity,
			CategoryID:  item.CategoryID,
			MediaID:     item.MediaID,
			URL:         url,
		},
	}, nil
}
