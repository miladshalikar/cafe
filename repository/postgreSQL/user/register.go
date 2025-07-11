package userpostgresql

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
)

func (u *UserDB) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {

	query := `INSERT INTO users (first_name, last_name, email, phone_number, password)
				VALUES ($1, $2, $3, $4, $5)
				RETURNING id`

	err := u.conn.QueryRowContext(ctx, query, user.FirstName, user.LastName,
		user.Email, user.PhoneNumber, user.GetPassword()).Scan(&user.ID)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
