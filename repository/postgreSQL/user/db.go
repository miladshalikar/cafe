package userpostgresql

import (
	"context"
	"database/sql"
	"errors"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

type UserDB struct {
	conn *sql.DB
}

func New(c *sql.DB) *UserDB {
	return &UserDB{conn: c}
}

func (u *UserDB) PhoneNumberExistInDB(ctx context.Context, phoneNumber string) (bool, error) {
	const op = "userpostgresql.PhoneNumberExistInDB"

	query := `SELECT * FROM users WHERE phone_number = $1`

	row := u.conn.QueryRowContext(ctx, query, phoneNumber)

	_, err := scanUser(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected).
			WithMeta(map[string]interface{}{"phoneNumber": phoneNumber})
	}
	return true, nil
}

func (u *UserDB) EmailExistInDB(ctx context.Context, email string) (bool, error) {
	const op = "userpostgresql.EmailExistInDB"

	query := `SELECT * FROM users WHERE email = $1`

	row := u.conn.QueryRowContext(ctx, query, email)

	_, err := scanUser(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected).
			WithMeta(map[string]interface{}{"email": email})
	}

	return true, nil
}
