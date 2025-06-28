package userpostgresql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/miladshalikar/cafe/entity"
)

func (u *UserDB) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	query := `select id, first_name, last_name, email, phone_number, created_at, updated_at from users where id = $1 AND deleted_at IS NULL`

	var user entity.User
	err := u.conn.QueryRowContext(ctx, query, id).Scan(
		&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, errors.New("user not found")
		}
		return entity.User{}, err
	}
	return user, nil
}
