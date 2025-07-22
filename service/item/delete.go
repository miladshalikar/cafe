package itemservice

import (
	"context"
	itemparam "github.com/miladshalikar/cafe/param/item"
	mediaparam "github.com/miladshalikar/cafe/param/media"
)

func (s Service) DeleteItem(ctx context.Context, req itemparam.DeleteItemRequest) (itemparam.DeleteItemResponse, error) {

	item, mErr := s.repo.GetItemByID(ctx, req.ID)
	if mErr != nil {
		return itemparam.DeleteItemResponse{}, mErr
	}

	if cErr := s.repo.DeleteItem(ctx, req.ID); cErr != nil {
		return itemparam.DeleteItemResponse{}, cErr
	}

	if _, dErr := s.client.DeleteMedia(ctx, mediaparam.DeleteMediaRequest{ID: item.MediaID}); dErr != nil {

		_ = s.repo.UndoDeleteItem(ctx, req.ID)

		return itemparam.DeleteItemResponse{}, dErr
	}

	return itemparam.DeleteItemResponse{}, nil
}
