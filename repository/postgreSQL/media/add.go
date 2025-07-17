package mediapostgresql

import (
	"context"
	"fmt"
	"github.com/miladshalikar/cafe/entity"
)

func (d *DB) AddMedia(ctx context.Context, media entity.Media) (entity.Media, error) {

	query := `INSERT INTO media (file_name,size,path,mime_type,is_private,bucket) 
				VALUES ($1,$2,$3,$4,$5,$6)
				RETURNING *`

	row := d.conn.QueryRowContext(ctx, query, media.FileName, media.Size, media.Path, media.MimeType, media.IsPrivate, media.Bucket)

	addedMedia, err := scanMedia(row)
	if err != nil {
		return entity.Media{}, fmt.Errorf("failed to scan category: %w", err)
	}

	return addedMedia, nil

}
