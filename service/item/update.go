package itemservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	itemparam "github.com/miladshalikar/cafe/param/item"
	mediaparam "github.com/miladshalikar/cafe/param/media"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (s Service) UpdateItem(ctx context.Context, req itemparam.UpdateItemRequest) (itemparam.UpdateItemResponse, error) {
	const op = "itemservice.UpdateCategory"

	currentItem, cErr := s.repo.GetItemByID(ctx, req.ID)
	if cErr != nil {
		return itemparam.UpdateItemResponse{}, richerror.New(op).WithWarpError(cErr)
	}

	if currentItem.MediaID != req.MediaID && req.MediaID != 0 {

		if _, sErr := s.client.DeleteMedia(ctx, mediaparam.DeleteMediaRequest{ID: currentItem.MediaID}); sErr != nil {
			return itemparam.UpdateItemResponse{}, richerror.New(op).WithWarpError(sErr)
		}
	}

	item := entity.Item{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
		CategoryID:  req.CategoryID,
	}
	if req.MediaID == 0 {
		item.MediaID = currentItem.MediaID
	} else {
		item.MediaID = req.MediaID
	}

	if rErr := s.repo.UpdateItem(ctx, item); rErr != nil {
		return itemparam.UpdateItemResponse{}, richerror.New(op).WithWarpError(rErr)
	}
	mediaURL, uErr := s.client.GetURLMedia(ctx, mediaparam.GetURLRequest{ID: item.MediaID})
	if uErr != nil {
		return itemparam.UpdateItemResponse{}, richerror.New(op).WithWarpError(uErr)
	}

	return itemparam.UpdateItemResponse{
		ItemInfo: itemparam.ItemInfo{
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			Price:       item.Price,
			Quantity:    item.Quantity,
			CategoryID:  item.CategoryID,
			MediaID:     item.MediaID,
			URL:         mediaURL.URL,
		},
	}, nil
}
