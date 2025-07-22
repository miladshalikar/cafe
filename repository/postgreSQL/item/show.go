package itempostgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/miladshalikar/cafe/entity"
)

func (d *DB) GetItemByID(ctx context.Context, id uint) (entity.Item, error) {

	query := `SELECT * FROM items WHERE id = $1 AND deleted_at IS NULL`

	row := d.conn.QueryRowContext(ctx, query, id)

	item, err := scanItem(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Item{}, fmt.Errorf("item with ID %d not found", id)
		}
		return entity.Item{}, fmt.Errorf("failed to scan item: %w", err)
	}

	return item, nil

}

func (d *DB) CheckItemIsExistByID(ctx context.Context, id uint) (bool, error) {

	query := `SELECT * FROM items WHERE id = $1 AND deleted_at IS NULL`

	row := d.conn.QueryRowContext(ctx, query, id)

	_, err := scanItem(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("failed to scan item: %w", err)
	}

	return true, nil
}
