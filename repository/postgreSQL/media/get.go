package mediapostgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/miladshalikar/cafe/entity"
)

func (d *DB) GetMediaByID(ctx context.Context, id uint) (entity.Media, error) {

	query := `SELECT * FROM media WHERE id = $1`

	row := d.conn.QueryRowContext(ctx, query, id)

	media, err := scanMedia(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Media{}, fmt.Errorf("media with ID %d not found", id)
		}
		return entity.Media{}, fmt.Errorf("failed to scan media: %w", err)
	}

	return media, nil
}

func (d *DB) CheckMediaIsExist(ctx context.Context, id uint) (bool, error) {

	query := `SELECT * FROM media WHERE id = $1`

	row := d.conn.QueryRowContext(ctx, query, id)

	_, err := scanMedia(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("failed to scan media: %w", err)
	}

	return true, nil
}
