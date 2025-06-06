package userpostgresql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/miladshalikar/cafe/entity"
)

func (u *UserDB) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {

	query := `INSERT INTO users (first_name, last_name, email, phone_number, password)
				VALUES ($1, $2, $3, $4, $5)
				RETURNING id`

	err := u.conn.QueryRowContext(ctx, query, user.FirstName, user.LastName,
		user.Email, user.PhoneNumber, user.GetPassword()).Scan(&user.Id)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u *UserDB) PhoneNumberExistInDB(ctx context.Context, phoneNumber string) (bool, error) {
	query := `SELECT * FROM users WHERE phone_number = $1`
	var exists int
	err := u.conn.QueryRowContext(ctx, query, phoneNumber).Scan(&exists)
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
	var exists int
	err := u.conn.QueryRowContext(ctx, query, email).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
