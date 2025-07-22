package itemservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	itemparam "github.com/miladshalikar/cafe/param/item"
)

func (s Service) AddNewItem(ctx context.Context, req itemparam.AddNewItemRequest) (itemparam.AddNewItemResponse, error) {

	iItem := entity.Item{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
		MediaID:     req.MediaID,
	}

	item, iErr := s.repo.AddNewItem(ctx, iItem)
	if iErr != nil {
		return itemparam.AddNewItemResponse{}, iErr
	}

	return itemparam.AddNewItemResponse{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		Price:       item.Price,
		CategoryID:  item.CategoryID,
		MediaID:     item.MediaID,
	}, nil
}
