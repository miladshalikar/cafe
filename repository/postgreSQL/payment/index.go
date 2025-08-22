package paymentpostgresql

import (
	"context"
	"fmt"
	"github.com/miladshalikar/cafe/entity"
	commonparam "github.com/miladshalikar/cafe/param/common"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (d *DB) GetTotalCountPaymentsByOrderID(ctx context.Context, orderID uint) (uint, error) {
	const op = "paymentpostgresql.GetTotalCountPaymentsByOrderID"

	query := `SELECT COUNT(*) FROM payments WHERE deleted_at IS NULL`

	var count uint

	if err := d.conn.QueryRowContext(ctx, query).Scan(&count); err != nil {
		return 0, richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).
			WithKind(richerror.KindUnexpected)
	}
	return count, nil
}

func (d *DB) GetPaymentsByOrderIDWithPagination(
	ctx context.Context,
	pagination commonparam.PaginationRequest,
	orderID uint,
) ([]entity.Payment, error) {
	const op = "paymentpostgresql.GetPaymentsByOrderIDWithPagination"

	argIndex := 1

	query := `SELECT * FROM payments WHERE deleted_at IS NULL`

	query += " ORDER BY id ASC"
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)

	rows, err := d.conn.QueryContext(ctx, query, pagination.GetPageSize(), pagination.GetOffset())
	if err != nil {
		return nil, richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).
			WithKind(richerror.KindUnexpected)
	}
	defer rows.Close()

	var payments []entity.Payment

	for rows.Next() {
		payment, pErr := scanPayment(rows)
		if pErr != nil {
			return nil, richerror.New(op).
				WithWarpError(err).
				WithMessage(errmsg.ErrorMsgCantScanQueryResult).
				WithKind(richerror.KindUnexpected)
		}
		payments = append(payments, payment)
	}
	if rErr := rows.Err(); rErr != nil {
		return nil, richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}

	return payments, nil
}
