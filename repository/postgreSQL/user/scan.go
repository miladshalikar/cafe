package userpostgresql

import (
	"github.com/miladshalikar/cafe/entity"
)

type Scanner interface {
	Scan(dest ...any) error
}

func scanUser(scanner Scanner) (entity.User, error) {

	var user entity.User
	var pass string
	var s bool

	err := scanner.Scan(&user.Id, &user.FirstName,
		&user.LastName, &user.Email, &user.PhoneNumber, &pass, &s, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	user.SetPassword(pass)

	return user, err
}
