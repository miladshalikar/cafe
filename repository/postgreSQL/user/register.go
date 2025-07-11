package userpostgresql

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
)

func (u *UserDB) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {

	query := `INSERT INTO users (first_name, last_name, email, phone_number, password)
				VALUES ($1, $2, $3, $4, $5)
				RETURNING *`

	row := u.conn.QueryRowContext(ctx, query, user.FirstName, user.LastName,
		user.Email, user.PhoneNumber, user.GetPassword())

	registeredUser, err := scanUser(row)

	if err != nil {
		return entity.User{}, err
	}
	return registeredUser, nil
}
