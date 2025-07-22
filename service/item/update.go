package itemservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	itemparam "github.com/miladshalikar/cafe/param/item"
	mediaparam "github.com/miladshalikar/cafe/param/media"
)

func (s Service) UpdateItem(ctx context.Context, req itemparam.UpdateItemRequest) (itemparam.UpdateItemResponse, error) {

	currentItem, cErr := s.repo.GetItemByID(ctx, req.ID)
	if cErr != nil {
		return itemparam.UpdateItemResponse{}, cErr
	}

	if currentItem.MediaID != req.MediaID {

		if _, sErr := s.client.DeleteMedia(ctx, mediaparam.DeleteMediaRequest{ID: currentItem.MediaID}); sErr != nil {
			return itemparam.UpdateItemResponse{}, sErr
		}
	}

	item := entity.Item{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
		MediaID:     req.MediaID,
	}

	if rErr := s.repo.UpdateItem(ctx, item); rErr != nil {
		return itemparam.UpdateItemResponse{}, rErr
	}
	mediaURL, uErr := s.client.GetURLMedia(ctx, mediaparam.GetURLRequest{ID: item.MediaID})
	if uErr != nil {
		return itemparam.UpdateItemResponse{}, uErr
	}

	return itemparam.UpdateItemResponse{
		ItemInfo: itemparam.ItemInfo{
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			Price:       item.Price,
			CategoryID:  item.CategoryID,
			MediaID:     item.MediaID,
			URL:         mediaURL.URL,
		},
	}, nil
}
