package itempostgresql

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
)

func (d *DB) AddNewItem(ctx context.Context, item entity.Item) (entity.Item, error) {
	query := `INSERT INTO items (title, description, price, category_id, media_id) 
				VALUES ($1, $2, $3, $4, $5) 
				RETURNING *`

	row := d.conn.QueryRowContext(ctx, query, item.Title, item.Description, item.Price, item.CategoryID, item.MediaID)

	addedItem, err := scanItem(row)
	if err != nil {
		return entity.Item{}, err
	}
	return addedItem, nil
}
