package paymentservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	commonparam "github.com/miladshalikar/cafe/param/common"
)

type Service struct {
	repo Repository
}

type Repository interface {
	CreatePayment(ctx context.Context, payment entity.Payment) (entity.Payment, error)
	SimulatePaymentResult(ctx context.Context, paymentID uint, status entity.PaymentStatus) error
	GetPaymentByID(ctx context.Context, paymentID uint) (entity.Payment, error)
	GetPaymentsByOrderIDWithPagination(ctx context.Context, pagination commonparam.PaginationRequest, orderID uint) ([]entity.Payment, error)
	GetTotalCountPaymentsByOrderID(ctx context.Context, orderID uint) (uint, error)
	RefundPayment(ctx context.Context, paymentID uint, status entity.PaymentStatus) error
}

func New(r Repository) Service {
	return Service{r}
}
