package itemservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	itemparam "github.com/miladshalikar/cafe/param/item"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (s Service) AddNewItem(ctx context.Context, req itemparam.AddNewItemRequest) (itemparam.AddNewItemResponse, error) {
	const op = "itemservice.AddNewItem"

	iItem := entity.Item{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
		CategoryID:  req.CategoryID,
		MediaID:     req.MediaID,
	}

	item, iErr := s.repo.AddNewItem(ctx, iItem)
	if iErr != nil {
		return itemparam.AddNewItemResponse{}, richerror.New(op).WithWarpError(iErr)
	}

	return itemparam.AddNewItemResponse{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		Price:       item.Price,
		Quantity:    item.Quantity,
		CategoryID:  item.CategoryID,
		MediaID:     item.MediaID,
	}, nil
}
