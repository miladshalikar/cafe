package authservice

import (
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
	CreateUser(entity.User) (entity.User, error)
}

func New(r Repository, t Tokens) Service {
	return Service{r, t}
}
