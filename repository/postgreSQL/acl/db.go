package aclpostgresql

import "database/sql"

type DB struct {
	conn *sql.DB
}

func New(c *sql.DB) *DB {
	return &DB{conn: c}
}
