package categorypostgresql

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
)

func (d *DB) AddNewCategory(ctx context.Context, category entity.Category) (entity.Category, error) {

	query := `INSERT INTO categories (title, media_id)
				VALUES ($1, $2)
				RETURNING *`

	row := d.conn.QueryRowContext(ctx, query, category.Title, category.MediaID)

	addedCategory, err := scanCategory(row)

	if err != nil {
		return entity.Category{}, err
	}
	return addedCategory, nil
}
