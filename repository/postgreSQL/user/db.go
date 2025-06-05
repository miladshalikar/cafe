package userpostgresql

import "database/sql"

type UserDB struct {
	conn *sql.DB
}

func New(c *sql.DB) *UserDB {
	return &UserDB{conn: c}
}
