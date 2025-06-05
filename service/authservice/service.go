package authservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
)

type Service struct {
	Repo   Repository
	Tokens Tokens
}

type Tokens interface {
	CreateAccessToken(adminID uint) (string, error)
	CreateRefreshToken(adminID uint) (string, error)
}

type Repository interface {
	CreateUser(context.Context, entity.User) (entity.User, error)
}

func New(r Repository, t Tokens) Service {
	return Service{r, t}
}
