package categorypostgresql

import (
	"context"
	"fmt"
	"github.com/miladshalikar/cafe/entity"
)

func (d *DB) GetTotalCountCategory(ctx context.Context) (uint, error) {

	query := `SELECT * FROM categories`

	var count uint

	if err := d.conn.QueryRowContext(ctx, query).Scan(&count); err != nil {
		return 0, fmt.Errorf("something went wrong: %w", err)
	}
	return count, nil

}

func (d *DB) GetCategoriesWithPagination(ctx context.Context, pageSize, offset uint) ([]entity.Category, error) {

	query := `SELECT * FROM categories LIMIT $1 OFFSET $2`

	rows, err := d.conn.QueryContext(ctx, query, pageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("query execution failed: %w", err)
	}
	defer rows.Close()

	var categories []entity.Category

	for rows.Next() {
		category, cErr := scanCategory(rows)
		if cErr != nil {
			return nil, fmt.Errorf("scanning category failed: %w", err)
		}
		categories = append(categories, category)
	}
	if rErr := rows.Err(); rErr != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return categories, nil
}
