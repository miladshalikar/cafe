package itempostgresql

import (
	"context"
	"fmt"
	"github.com/miladshalikar/cafe/entity"
)

func (d *DB) GetTotalCountItem(ctx context.Context, search string) (uint, error) {

	query := `SELECT COUNT(*) FROM items WHERE deleted_at IS NULL`

	var count uint
	var args []interface{}

	if search != "" {
		query += " AND title ILIKE $1"
		args = append(args, "%"+search+"%")
	}

	if err := d.conn.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, fmt.Errorf("something went wrong: %w", err)
	}
	return count, nil
}

func (d *DB) GetItemsWithPagination(ctx context.Context, pageSize uint, offset uint, search string) ([]entity.Item, error) {

	query := `SELECT * FROM items WHERE deleted_at IS NULL`

	var args []interface{}
	argIndex := 1

	if search != "" {
		query += fmt.Sprintf(" AND title ILIKE $%d", argIndex)
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

	var items []entity.Item

	for rows.Next() {
		item, iErr := scanItem(rows)
		if iErr != nil {
			return nil, fmt.Errorf("scanning item failed: %w", err)
		}
		items = append(items, item)
	}
	if rErr := rows.Err(); rErr != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return items, nil
}
