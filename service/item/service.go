package itemservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	mediaparam "github.com/miladshalikar/cafe/param/media"
)

type Service struct {
	repo   Repository
	client Client
	cache  Cache
}

type Repository interface {
	AddNewItem(ctx context.Context, item entity.Item) (entity.Item, error)
	DeleteItem(ctx context.Context, id uint) error
	UndoDeleteItem(ctx context.Context, id uint) error
	GetItemByID(ctx context.Context, id uint) (entity.Item, error)
	UpdateItem(ctx context.Context, item entity.Item) error
	GetTotalCountItem(ctx context.Context, search string) (uint, error)
	GetItemsWithPagination(ctx context.Context, pageSize, offset uint, search string) ([]entity.Item, error)
}

type Client interface {
	UploadMedia(ctx context.Context, req mediaparam.UploadMediaRequest) (mediaparam.UploadMediaResponse, error)
	GetURLMedia(ctx context.Context, req mediaparam.GetURLRequest) (mediaparam.GetURLResponse, error)
	DeleteMedia(ctx context.Context, req mediaparam.DeleteMediaRequest) (mediaparam.DeleteMediaResponse, error)
}

type Cache interface {
	SetMediaURLByMediaID(ctx context.Context, mediaID uint, url string) error
	GetMediaURLByMediaID(ctx context.Context, mediaID uint) (string, error)
	MGetMediaURLs(ctx context.Context, mediaIDs []uint) (map[uint]string, error)
}

func New(r Repository, c Client, ca Cache) Service {
	return Service{repo: r, client: c, cache: ca}
}
