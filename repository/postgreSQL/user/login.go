package userpostgresql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/miladshalikar/cafe/entity"
)

func (u *UserDB) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {

	query := `select id, first_name, last_name, email, phone_number, password from users where email = $1`

	var user entity.User
	var p string
	err := u.conn.QueryRowContext(ctx, query, email).Scan(
		&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber, &p)
	user.SetPassword(p)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, errors.New("user not found")
		}
		return entity.User{}, err
	}
	return user, nil
}
