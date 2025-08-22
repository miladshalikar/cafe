package paymentpostgresql

import (
	"github.com/miladshalikar/cafe/entity"
	postgresql "github.com/miladshalikar/cafe/repository"
)

func scanPayment(scanner postgresql.Scanner) (entity.Payment, error) {

	var p entity.Payment

	err := scanner.Scan(&p.ID, &p.OrderID, &p.Amount, &p.Status, &p.Method,
		&p.CreatedAt, &p.UpdatedAt, &p.DeletedAt)

	return p, err
}
