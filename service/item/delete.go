package itemservice

import (
	"context"
	itemparam "github.com/miladshalikar/cafe/param/item"
	mediaparam "github.com/miladshalikar/cafe/param/media"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (s Service) DeleteItem(ctx context.Context, req itemparam.DeleteItemRequest) (itemparam.DeleteItemResponse, error) {
	const op = "itemservice.DeleteItem"

	item, mErr := s.repo.GetItemByID(ctx, req.ID)
	if mErr != nil {
		return itemparam.DeleteItemResponse{}, richerror.New(op).WithWarpError(mErr)
	}

	//todo
	if _, dErr := s.client.DeleteMedia(ctx, mediaparam.DeleteMediaRequest{ID: item.MediaID}); dErr != nil {
		return itemparam.DeleteItemResponse{}, richerror.New(op).WithWarpError(dErr).WithMeta(map[string]interface{}{"media_id": item.MediaID})
	}

	if cErr := s.repo.DeleteItem(ctx, req.ID); cErr != nil {
		return itemparam.DeleteItemResponse{}, richerror.New(op).WithWarpError(cErr)
	}

	return itemparam.DeleteItemResponse{}, nil
}
