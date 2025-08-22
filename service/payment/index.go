package paymentservice

import (
	"context"
	commonparam "github.com/miladshalikar/cafe/param/common"
	paymentparam "github.com/miladshalikar/cafe/param/payment"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (s Service) GetPaymentsByOrderID(ctx context.Context, req paymentparam.GetPaymentsByOrderIDRequest) (paymentparam.GetPaymentsByOrderIDResponse, error) {
	const op = "paymentservice.GetPaymentsByOrderID"

	total, tErr := s.repo.GetTotalCountPaymentsByOrderID(ctx, req.OrderID)
	if tErr != nil {
		return paymentparam.GetPaymentsByOrderIDResponse{}, richerror.New(op).WithWarpError(tErr)
	}

	payments, pErr := s.repo.GetPaymentsByOrderIDWithPagination(ctx, req.Pagination, req.OrderID)
	if pErr != nil {
		return paymentparam.GetPaymentsByOrderIDResponse{}, richerror.New(op).WithWarpError(pErr)
	}

	var paymentsInfo []paymentparam.PaymentInfo
	for _, payment := range payments {
		paymentsInfo = append(paymentsInfo, paymentparam.PaymentInfo{
			ID:      payment.ID,
			OrderID: payment.OrderID,
			Amount:  payment.Amount,
			Status:  payment.Status,
			Method:  payment.Method,
		})
	}

	return paymentparam.GetPaymentsByOrderIDResponse{
		Pagination: commonparam.PaginationResponse{
			PageSize:   req.Pagination.GetPageSize(),
			PageNumber: req.Pagination.GetPageNumber(),
			Total:      total,
		},
		Payments: paymentsInfo,
	}, nil
}
