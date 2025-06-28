package userauthservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
)

type Service struct {
	repo   Repository
	tokens Tokens
}

type Tokens interface {
	CreateAccessToken(adminID uint) (string, error)
	CreateRefreshToken(adminID uint) (string, error)
}

type Repository interface {
	CreateUser(context.Context, entity.User) (entity.User, error)
	GetUserByEmail(context.Context, string) (entity.User, error)
}

func New(r Repository, t Tokens) Service {
	return Service{r, t}
}
