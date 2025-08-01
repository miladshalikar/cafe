package itempostgresql

import (
	"context"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (d *DB) DeleteItem(ctx context.Context, id uint) error {
	const op = "itempostgresql.DeleteItem"

	query := `UPDATE items SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`

	result, err := d.conn.ExecContext(ctx, query, id)
	if err != nil {
		return richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).
			WithKind(richerror.KindUnexpected)
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).
			WithKind(richerror.KindUnexpected)
	}

	if rowAffected == 0 {
		return richerror.New(op).
			WithMessage(errmsg.ErrorMsgNotFound).
			WithKind(richerror.KindNotFound).
			WithMeta(map[string]interface{}{"id": id})
	}

	return nil
}

func (d *DB) UndoDeleteItem(ctx context.Context, id uint) error {
	//todo
	query := `UPDATE items SET deleted_at = NULL WHERE id = $1`
	_, err := d.conn.ExecContext(ctx, query, id)
	return err
}
