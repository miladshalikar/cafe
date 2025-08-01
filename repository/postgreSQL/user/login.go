package userpostgresql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/miladshalikar/cafe/entity"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (u *UserDB) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	const op = "userpostgresql.GetUserByEmail"

	query := `SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL`

	row := u.conn.QueryRowContext(ctx, query, email)

	user, err := scanUser(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, richerror.New(op).
				WithWarpError(err).
				WithMessage(errmsg.ErrorMsgNotFound).
				WithKind(richerror.KindNotFound)
		}
		return entity.User{}, richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}
	return user, nil
}
