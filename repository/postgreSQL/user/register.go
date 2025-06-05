package userpostgresql

import (
	"database/sql"
	"github.com/miladshalikar/cafe/entity"
)

type UserDB struct {
	conn *sql.DB
}

func New(con *sql.DB) UserDB {
	return UserDB{conn: con}
}

func (u *UserDB) CreateUser(user entity.User) (entity.User, error) {
	panic("implement me")
}
