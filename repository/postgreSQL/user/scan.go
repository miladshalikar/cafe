package userpostgresql

import (
	"github.com/miladshalikar/cafe/entity"
	postgresql "github.com/miladshalikar/cafe/repository"
)

func scanUser(scanner postgresql.Scanner) (entity.User, error) {

	var user entity.User
	var pass string
	var s bool

	err := scanner.Scan(&user.Id, &user.FirstName,
		&user.LastName, &user.Email, &user.PhoneNumber, &pass, &s, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	user.SetPassword(pass)

	return user, err
}
