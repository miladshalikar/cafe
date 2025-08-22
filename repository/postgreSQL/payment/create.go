package paymentpostgresql

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (d *DB) CreatePayment(ctx context.Context, payment entity.Payment) (entity.Payment, error) {
	const op = "paymentpostgresql.CreatePayment"

	query := `INSERT INTO payments (order_id, amount, status, method) 
				VALUES ($1, $2, $3, $4) 
				RETURNING *`

	row := d.conn.QueryRowContext(ctx, query, payment.OrderID, payment.Amount, payment.Status, payment.Method)

	addedPayment, err := scanPayment(row)
	if err != nil {
		return entity.Payment{}, richerror.New(op).WithWarpError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}
	return addedPayment, nil
}
