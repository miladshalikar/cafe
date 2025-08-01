package itempostgresql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/miladshalikar/cafe/entity"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (d *DB) GetItemByID(ctx context.Context, id uint) (entity.Item, error) {
	const op = "itempostgresql.GetItemByID"

	query := `SELECT * FROM items WHERE id = $1 AND deleted_at IS NULL`

	row := d.conn.QueryRowContext(ctx, query, id)

	item, err := scanItem(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Item{}, richerror.New(op).
				WithWarpError(err).
				WithMessage(errmsg.ErrorMsgNotFound).
				WithKind(richerror.KindNotFound)
		}
		return entity.Item{}, richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}

	return item, nil
}

func (d *DB) CheckItemIsExistByID(ctx context.Context, id uint) (bool, error) {
	const op = "itempostgresql.CheckItemIsExistByID"

	query := `SELECT * FROM items WHERE id = $1 AND deleted_at IS NULL`

	row := d.conn.QueryRowContext(ctx, query, id)

	_, err := scanItem(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected).
			WithMeta(map[string]interface{}{"id": id})
	}

	return true, nil
}
