package userpostgresql

import (
	"context"
	"database/sql"
	"errors"
)

type UserDB struct {
	conn *sql.DB
}

func New(c *sql.DB) *UserDB {
	return &UserDB{conn: c}
}

func (u *UserDB) PhoneNumberExistInDB(ctx context.Context, phoneNumber string) (bool, error) {
	query := `SELECT * FROM users WHERE phone_number = $1`

	row := u.conn.QueryRowContext(ctx, query, phoneNumber)

	_, err := scanUser(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (u *UserDB) EmailExistInDB(ctx context.Context, email string) (bool, error) {
	query := `SELECT * FROM users WHERE email = $1`

	row := u.conn.QueryRowContext(ctx, query, email)

	_, err := scanUser(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
