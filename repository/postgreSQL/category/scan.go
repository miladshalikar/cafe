package categorypostgresql

import (
	"github.com/miladshalikar/cafe/entity"
	postgresql "github.com/miladshalikar/cafe/repository"
)

func scanCategory(scanner postgresql.Scanner) (entity.Category, error) {
	var category entity.Category

	err := scanner.Scan(&category.ID, &category.Title, &category.MediaID, &category.CreatedAt, &category.UpdatedAt, &category.DeletedAt)

	return category, err
}
