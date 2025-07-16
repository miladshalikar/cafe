package categorypostgresql

import (
	"context"
	"fmt"
	"github.com/miladshalikar/cafe/entity"
)

func (d *DB) GetTotalCountCategory(ctx context.Context, search string) (uint, error) {

	query := `SELECT COUNT(*) FROM categories`

	var count uint
	var args []interface{}

	if search != "" {
		query += " WHERE title ILIKE $1"
		args = append(args, "%"+search+"%")
	}

	if err := d.conn.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, fmt.Errorf("something went wrong: %w", err)
	}
	return count, nil

}

func (d *DB) GetCategoriesWithPagination(ctx context.Context, pageSize, offset uint, search string) ([]entity.Category, error) {

	query := `SELECT * FROM categories`
	var args []interface{}
	argIndex := 1

	if search != "" {
		query += fmt.Sprintf(" WHERE title ILIKE $%d", argIndex)
		args = append(args, "%"+search+"%")
		argIndex++
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, pageSize, offset)

	rows, err := d.conn.QueryContext(ctx, query, args...)
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
