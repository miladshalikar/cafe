package categorypostgresql

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (d *DB) AddNewCategory(ctx context.Context, category entity.Category) (entity.Category, error) {
	const op = "categorypostgresql.AddNewCategory"

	query := `INSERT INT categories (title, media_id)
				VALUES ($1, $2)
				RETURNING *`

	row := d.conn.QueryRowContext(ctx, query, category.Title, category.MediaID)
	addedCategory, err := scanCategory(row)

	if err != nil {
		return entity.Category{}, richerror.New(op).WithWarpError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}
	return addedCategory, nil
}
