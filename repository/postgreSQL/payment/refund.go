package paymentpostgresql

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (d *DB) RefundPayment(ctx context.Context, paymentID uint, status entity.PaymentStatus) error {
	const op = "paymentpostgresql.RefundPayment"

	query := `UPDATE payments SET 
                    status = $1
                    WHERE id = $2 AND deleted_at IS NULL`

	result, err := d.conn.ExecContext(ctx, query, status, paymentID)
	if err != nil {
		return richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).
			WithKind(richerror.KindUnexpected)
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).
			WithKind(richerror.KindUnexpected)
	}

	if rowAffected == 0 {
		return richerror.New(op).
			WithMessage(errmsg.ErrorMsgNotFound).
			WithKind(richerror.KindNotFound).
			WithMeta(map[string]interface{}{"paymentID": paymentID, "status": status})
	}

	return nil
}
