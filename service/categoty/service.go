package categoryservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	mediaparam "github.com/miladshalikar/cafe/param/media"
)

type Service struct {
	repo   Repository
	Client Client
}

type Repository interface {
	AddNewCategory(ctx context.Context, category entity.Category) (entity.Category, error)
	DeleteCategory(ctx context.Context, id uint) error
	GetCategoryByID(ctx context.Context, id uint) (entity.Category, error)
	UpdateCategory(ctx context.Context, category entity.Category) error
	GetTotalCountCategory(ctx context.Context, search string) (uint, error)
	GetCategoriesWithPagination(ctx context.Context, pageSize, offset uint, search string) ([]entity.Category, error)
}

type Client interface {
	UploadMedia(ctx context.Context, req mediaparam.UploadMediaRequest) (mediaparam.UploadMediaResponse, error)
	GetURLMedia(ctx context.Context, req mediaparam.GetURLRequest) (mediaparam.GetURLResponse, error)
	DeleteMedia(ctx context.Context, req mediaparam.DeleteMediaRequest) (mediaparam.DeleteMediaResponse, error)
}

func New(r Repository, c Client) Service {
	return Service{repo: r, Client: c}
}
