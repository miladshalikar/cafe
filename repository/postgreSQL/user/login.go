package userpostgresql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/miladshalikar/cafe/entity"
)

func (u *UserDB) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {

	query := `select * from users where email = $1`

	row := u.conn.QueryRowContext(ctx, query, email)

	user, err := scanUser(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, errors.New("user not found")
		}
		return entity.User{}, err
	}
	return user, nil
}
