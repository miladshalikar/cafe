package categorypostgresql

import (
	"database/sql"
	"github.com/miladshalikar/cafe/entity"
	postgresql "github.com/miladshalikar/cafe/repository"
)

func scanCategory(scanner postgresql.Scanner) (entity.Category, error) {
	var category entity.Category
	var mediaID sql.NullInt64

	err := scanner.Scan(&category.ID, &category.Title, &mediaID, &category.CreatedAt, &category.UpdatedAt, &category.DeletedAt)

	if mediaID.Valid {
		category.MediaID = uint(mediaID.Int64)
	} else {
		category.MediaID = 0
	}

	return category, err
}
