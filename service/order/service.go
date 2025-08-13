package orderservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	commonparam "github.com/miladshalikar/cafe/param/common"
	"time"
)

type Service struct {
	repo Repository
}

type Repository interface {
	AddNewOrder(ctx context.Context, order entity.Order) (entity.Order, error)
	UpdateOrder(ctx context.Context, order entity.Order) error
	UpdateOrderStatus(ctx context.Context, orderID uint, newStatus entity.OrderStatus) error
	CancelOrder(ctx context.Context, orderID uint) error
	DeleteOrder(ctx context.Context, orderID uint) error
	GetOrderByID(ctx context.Context, orderID uint) (entity.Order, error)
	GetOrdersByUserID(ctx context.Context, userID uint) ([]entity.Order, error)
	GetTotalCountOrderWithFilter(ctx context.Context, filter orderparam.FilterRequest) (uint, error)
	GetOrdersWithPaginationAndFilter(ctx context.Context, pagination commonparam.PaginationRequest, filter orderparam.FilterRequest) ([]entity.Order, error)
	GetDailyIncome(date time.Time) (float64, error)
	GetMonthlyIncome(date time.Time) (float64, error)
}
