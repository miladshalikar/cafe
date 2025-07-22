package itempostgresql

import (
	"context"
	"fmt"
)

func (d *DB) DeleteItem(ctx context.Context, id uint) error {

	query := `UPDATE items SET deleted_at = NOW() WHERE id = 1 AND deleted_at IS NOT NULL`

	result, err := d.conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return fmt.Errorf("no item found with id %d", id)
	}

	return nil
}

func (d *DB) undoDeleteItem(ctx context.Context, id uint) error {
	query := `UPDATE items SET deleted_at = NULL WHERE id = $1`
	_, err := d.conn.ExecContext(ctx, query, id)
	return err
}
