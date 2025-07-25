package categoryservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	commonparam "github.com/miladshalikar/cafe/param/common"
	mediaparam "github.com/miladshalikar/cafe/param/media"
)

type Service struct {
	repo   Repository
	client Client
	cache  Cache
}

type Repository interface {
	AddNewCategory(ctx context.Context, category entity.Category) (entity.Category, error)
	DeleteCategory(ctx context.Context, id uint) error
	UndoDeleteCategory(ctx context.Context, id uint) error
	GetCategoryByID(ctx context.Context, id uint) (entity.Category, error)
	UpdateCategory(ctx context.Context, category entity.Category) error
	GetTotalCountCategoryWithSearch(ctx context.Context, search commonparam.SearchRequest) (uint, error)
	GetCategoriesWithPaginationAndSearch(ctx context.Context, pagination commonparam.PaginationRequest, search commonparam.SearchRequest) ([]entity.Category, error)
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
