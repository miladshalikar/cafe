package categorypostgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/miladshalikar/cafe/entity"
)

func (d *DB) GetCategoryById(ctx context.Context, id uint) (entity.Category, error) {

	query := `SELECT * FROM categories WHERE id = $1`

	row := d.conn.QueryRowContext(ctx, query, id)

	category, err := scanCategory(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Category{}, fmt.Errorf("category with ID %d not found", id)
		}
		return entity.Category{}, fmt.Errorf("failed to scan category: %w", err)
	}

	return category, nil

}
