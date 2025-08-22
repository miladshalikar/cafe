package paymentservice

import (
	"context"
	paymentparam "github.com/miladshalikar/cafe/param/payment"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (s Service) GetPaymentByID(ctx context.Context, req paymentparam.ShowSinglePaymentRequest) (paymentparam.ShowSinglePaymentResponse, error) {
	const op = "paymentservice.GetPaymentByID"

	payment, err := s.repo.GetPaymentByID(ctx, req.ID)
	if err != nil {
		return paymentparam.ShowSinglePaymentResponse{}, richerror.New(op).WithWarpError(err)
	}
	return paymentparam.ShowSinglePaymentResponse{
		PaymentInfo: paymentparam.PaymentInfo{
			ID:      payment.ID,
			OrderID: payment.OrderID,
			Amount:  payment.Amount,
			Status:  payment.Status,
			Method:  payment.Method,
		},
	}, nil
}
