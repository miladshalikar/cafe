package categorypostgresql

import (
	"context"
	"fmt"
	"github.com/miladshalikar/cafe/entity"
)

func (d *DB) UpdateCategory(ctx context.Context, category entity.Category) error {

	query := `UPDATE categories SET title = $1, media_id = $2 WHERE id = $3 AND deleted_at IS NULL`

	result, err := d.conn.ExecContext(ctx, query, category.Title, category.MediaID, category.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no category found with id %d", category.ID)
	}

	return nil
}
