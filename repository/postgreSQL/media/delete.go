package mediapostgresql

import (
	"context"
	"fmt"
)

func (d *DB) DeleteMedia(ctx context.Context, id uint) error {

	query := `DELETE FROM media WHERE id=$1`

	result, err := d.conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no media found with id %d", id)
	}

	return nil
}
