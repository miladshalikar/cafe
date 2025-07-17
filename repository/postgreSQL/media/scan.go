package mediapostgresql

import (
	"github.com/miladshalikar/cafe/entity"
	postgresql "github.com/miladshalikar/cafe/repository"
)

func scanMedia(scanner postgresql.Scanner) (entity.Media, error) {

	var media entity.Media

	err := scanner.Scan(&media.ID, &media.FileName, &media.Size,
		&media.Path, &media.MimeType, &media.IsPrivate, &media.Bucket,
		&media.CreatedAt, &media.UpdatedAt, &media.DeletedAt)

	return media, err
}
