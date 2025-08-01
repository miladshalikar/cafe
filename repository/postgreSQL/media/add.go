package mediapostgresql

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (d *DB) AddMedia(ctx context.Context, media entity.Media) (entity.Media, error) {
	const op = "mediapostgresql.AddMedia"

	query := `INSERT INTO media (file_name,size,path,mime_type,is_private,bucket) 
				VALUES ($1,$2,$3,$4,$5,$6)
				RETURNING *`

	row := d.conn.QueryRowContext(ctx, query, media.FileName, media.Size, media.Path, media.MimeType, media.IsPrivate, media.Bucket)

	addedMedia, err := scanMedia(row)
	if err != nil {
		return entity.Media{}, richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).
			WithKind(richerror.KindUnexpected)
	}

	return addedMedia, nil

}
