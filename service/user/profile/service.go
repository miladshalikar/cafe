package userprofileservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
)

type Service struct {
	repo Repository
}

type Repository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
}

func New(r Repository) Service {
	return Service{r}
}
