package mediapostgresql

import (
	"context"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (d *DB) DeleteMedia(ctx context.Context, id uint) error {
	const op = "mediapostgresql.DeleteMedia"

	query := `UPDATE media
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL`

	result, err := d.conn.ExecContext(ctx, query, id)
	if err != nil {
		return richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).
			WithKind(richerror.KindUnexpected)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).
			WithKind(richerror.KindUnexpected)
	}

	if rowsAffected == 0 {
		return richerror.New(op).
			WithMessage(errmsg.ErrorMsgNotFound).
			WithKind(richerror.KindNotFound).
			WithMeta(map[string]interface{}{"id": id})
	}

	return nil
}
