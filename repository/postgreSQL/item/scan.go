package itempostgresql

import (
	"github.com/miladshalikar/cafe/entity"
	postgresql "github.com/miladshalikar/cafe/repository"
)

func scanItem(scanner postgresql.Scanner) (entity.Item, error) {

	var item entity.Item

	err := scanner.Scan(&item.ID, &item.Title, &item.Description, &item.Price, &item.CategoryID, &item.MediaID,
		&item.CreatedAt, &item.UpdatedAt, &item.DeletedAt)

	return item, err
}
