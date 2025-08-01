package categorypostgresql

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (d *DB) UpdateCategory(ctx context.Context, category entity.Category) error {
	const op = "categorypostgresql.UpdateCategory"

	query := `UPDATE categories SET title = $1, media_id = $2 WHERE id = $3 AND deleted_at IS NULL`

	result, err := d.conn.ExecContext(ctx, query, category.Title, category.MediaID, category.ID)
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
			WithMeta(map[string]interface{}{"category": category})
	}

	return nil
}
