package entity

import "time"

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"   // ایجاد شده ولی هنوز پرداخت نشده
	PaymentStatusCompleted PaymentStatus = "completed" // پرداخت موفق
	PaymentStatusFailed    PaymentStatus = "failed"    // خطا در پرداخت
	PaymentStatusRefunded  PaymentStatus = "refunded"  // بازگشت وجه
)

type Payment struct {
	ID        uint          `json:"id"`
	OrderID   uint          `json:"order_id"`
	Amount    float64       `json:"amount"`
	Status    PaymentStatus `json:"status"`
	Method    string        `json:"method"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	DeletedAt *time.Time    `json:"deleted_at"`
}
