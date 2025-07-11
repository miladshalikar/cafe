package userpostgresql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/miladshalikar/cafe/entity"
)

func (u *UserDB) GetUserByID(ctx context.Context, id int) (entity.User, error) {

	query := `select * from users where id = $1 AND deleted_at IS NULL`

	row := u.conn.QueryRowContext(ctx, query, id)

	user, err := scanUser(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, errors.New("user not found")
		}
		return entity.User{}, err
	}
	return user, nil
}
