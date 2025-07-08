package categorypostgresql

import (
	"context"
	"fmt"
)

func (d *DB) DeleteCategory(ctx context.Context, id uint) error {

	query := `DELETE FROM categories WHERE id = $1`

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
