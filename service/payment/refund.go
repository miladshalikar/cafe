package paymentservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	paymentparam "github.com/miladshalikar/cafe/param/payment"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (s Service) RefundPayment(ctx context.Context, req paymentparam.RefundPaymentRequest) (paymentparam.RefundPaymentResponse, error) {
	const op = "paymentservice.RefundPayment"

	err := s.repo.RefundPayment(ctx, req.PaymentID, entity.PaymentStatusRefunded)
	if err != nil {
		return paymentparam.RefundPaymentResponse{}, richerror.New(op).WithWarpError(err)
	}
	return paymentparam.RefundPaymentResponse{
		PaymentID: req.PaymentID,
		Status:    entity.PaymentStatusRefunded,
		Message:   "Payment has been successfully refunded.",
	}, nil
}
