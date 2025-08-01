package itempostgresql

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (d *DB) AddNewItem(ctx context.Context, item entity.Item) (entity.Item, error) {
	const op = "itempostgresql.AddNewItem"

	query := `INSERT INTO items (title, description, price, category_id, media_id) 
				VALUES ($1, $2, $3, $4, $5) 
				RETURNING *`

	row := d.conn.QueryRowContext(ctx, query, item.Title, item.Description, item.Price, item.CategoryID, item.MediaID)

	addedItem, err := scanItem(row)
	if err != nil {
		return entity.Item{}, richerror.New(op).WithWarpError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}
	return addedItem, nil
}
