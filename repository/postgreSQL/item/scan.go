package itempostgresql

import (
	"database/sql"
	"github.com/miladshalikar/cafe/entity"
	postgresql "github.com/miladshalikar/cafe/repository"
)

func scanItem(scanner postgresql.Scanner) (entity.Item, error) {

	var item entity.Item
	var mediaID sql.NullInt64

	err := scanner.Scan(&item.ID, &item.Title, &item.Description, &item.Price, &item.CategoryID, &mediaID,
		&item.CreatedAt, &item.UpdatedAt, &item.DeletedAt)

	if mediaID.Valid {
		item.MediaID = uint(mediaID.Int64)
	} else {
		item.MediaID = 0
	}

	return item, err
}
