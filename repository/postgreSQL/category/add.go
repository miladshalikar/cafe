package categorypostgresql

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
)

func (d *DB) AddNewCategory(ctx context.Context, category entity.Category) (entity.Category, error) {

	query := `INSERT INTO categories (title, logo)
				VALUES ($1, $2)
				RETURNING id`

	err := d.conn.QueryRowContext(ctx, query, category.Title, category.Logo).Scan(&category.Id)
	if err != nil {
		return entity.Category{}, err
	}
	return category, nil
}
