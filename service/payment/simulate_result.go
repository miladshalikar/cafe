package paymentservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	paymentparam "github.com/miladshalikar/cafe/param/payment"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (s Service) SimulatePaymentResult(ctx context.Context, req paymentparam.SimulatePaymentResultRequest) (paymentparam.SimulatePaymentResultResponse, error) {
	const op = "paymentservice.SimulatePaymentResult"

	status := entity.PaymentStatusFailed
	if req.Success {
		status = entity.PaymentStatusCompleted
	}

	err := s.repo.SimulatePaymentResult(ctx, req.PaymentID, status)
	if err != nil {
		return paymentparam.SimulatePaymentResultResponse{}, richerror.New(op).WithWarpError(err)
	}

	return paymentparam.SimulatePaymentResultResponse{}, nil
}
