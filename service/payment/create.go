package paymentservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	paymentparam "github.com/miladshalikar/cafe/param/payment"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (s Service) CreatePayment(ctx context.Context, req paymentparam.CreatePaymentRequest) (paymentparam.CreatePaymentResponse, error) {
	const op = "paymentservice.CreatePayment"

	payment := entity.Payment{
		OrderID: req.OrderID,
		Amount:  req.Amount,
		Method:  req.Method,
		Status:  entity.PaymentStatusPending,
	}

	createdPayment, err := s.repo.CreatePayment(ctx, payment)
	if err != nil {
		return paymentparam.CreatePaymentResponse{}, richerror.New(op).WithWarpError(err)
	}

	return paymentparam.CreatePaymentResponse{
		ID:      createdPayment.ID,
		OrderID: createdPayment.OrderID,
		Amount:  createdPayment.Amount,
		Status:  createdPayment.Status,
		Method:  createdPayment.Method,
	}, nil
}
