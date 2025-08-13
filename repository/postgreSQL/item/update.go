package itempostgresql

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (d *DB) UpdateItem(ctx context.Context, item entity.Item) error {
	const op = "itempostgresql.UpdateItem"

	query := `UPDATE items SET 
                 title = $1,
                 description = $2,
                 price = $3,
                 quantity = $4,
                 Category_id = $5,
                 media_id = $6 
             WHERE id = $7 AND deleted_at IS NULL`

	result, err := d.conn.ExecContext(ctx, query, item.Title, item.Description, item.Price, item.Quantity, item.CategoryID, item.MediaID, item.ID)
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
			WithMeta(map[string]interface{}{"item": item})
	}

	return nil
}
