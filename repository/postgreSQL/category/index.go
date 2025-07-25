package categorypostgresql

import (
	"context"
	"fmt"
	"github.com/miladshalikar/cafe/entity"
	commonparam "github.com/miladshalikar/cafe/param/common"
)

func (d *DB) GetTotalCountCategoryWithSearch(ctx context.Context, search commonparam.SearchRequest) (uint, error) {

	query := `SELECT COUNT(*) FROM categories WHERE deleted_at IS NULL`

	var count uint
	var args []interface{}

	if search.Search != "" {
		query += " AND title ILIKE $1"
		args = append(args, "%"+search.Search+"%")
	}

	if err := d.conn.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, fmt.Errorf("something went wrong: %w", err)
	}
	return count, nil

}

func (d *DB) GetCategoriesWithPaginationAndSearch(ctx context.Context, pagination commonparam.PaginationRequest, search commonparam.SearchRequest) ([]entity.Category, error) {

	query := `SELECT * FROM categories WHERE deleted_at IS NULL`
	var args []interface{}
	argIndex := 1

	if search.Search != "" {
		query += fmt.Sprintf(" AND title ILIKE $%d", argIndex)
		args = append(args, "%"+search.Search+"%")
		argIndex++
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, pagination.GetPageSize(), pagination.GetOffset())

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
