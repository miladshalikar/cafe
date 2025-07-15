package categoryservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
)

type Service struct {
	repo Repository
}

type Repository interface {
	AddNewCategory(ctx context.Context, category entity.Category) (entity.Category, error)
	DeleteCategory(ctx context.Context, id uint) error
	GetCategoryById(ctx context.Context, id uint) (entity.Category, error)
	UpdateCategory(ctx context.Context, category entity.Category) error
	GetTotalCountCategory(ctx context.Context) (uint, error)
	GetCategoriesWithPagination(ctx context.Context, pageSize, offset uint) ([]entity.Category, error)
}

func New(r Repository) Service {
	return Service{repo: r}
}
