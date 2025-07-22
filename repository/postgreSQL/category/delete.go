package categorypostgresql

import (
	"context"
	"fmt"
)

func (d *DB) DeleteCategory(ctx context.Context, id uint) error {

	query := `UPDATE categories
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL`

	result, err := d.conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no category found with id %d", id)
	}

	return nil
}

func (d *DB) UndoDeleteCategory(ctx context.Context, id uint) error {

	query := `UPDATE categories SET deleted_at = NULL WHERE id = $1`
	_, err := d.conn.ExecContext(ctx, query, id)
	return err
}
