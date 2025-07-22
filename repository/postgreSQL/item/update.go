package itempostgresql

import (
	"context"
	"fmt"
	"github.com/miladshalikar/cafe/entity"
)

func (d *DB) UpdateItem(ctx context.Context, item entity.Item) error {

	query := `UPDATE items SET 
                 title = $1,
                 description = $2,
                 price = $3,
                 Category_id = $4,
                 media_id = $5 
             WHERE id = $6 AND deleted_at IS NULL`

	result, err := d.conn.ExecContext(ctx, query, item.Title, item.Description, item.Price, item.CategoryID, item.MediaID, item.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no category found with id %d", item.ID)
	}

	return nil
}
