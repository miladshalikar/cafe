package entity

import "time"

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"    // در انتظار تأیید
	OrderStatusAccepted   OrderStatus = "accepted"   // تأیید ‌شده
	OrderStatusRejected   OrderStatus = "rejected"   // رد ‌شده
	OrderStatusProcessing OrderStatus = "processing" // در حال پردازش
	OrderStatusPaid       OrderStatus = "paid"       // پرداخت ‌شده
	OrderStatusShipped    OrderStatus = "shipped"    // ارسال‌ شده
	OrderStatusDelivered  OrderStatus = "delivered"  // تحویل داده ‌شده
	OrderStatusCanceled   OrderStatus = "canceled"   // لغو ‌شده
)

type Order struct {
	Id         uint        `json:"id"`
	UserId     uint        `json:"user_id"`
	Status     OrderStatus `json:"status"`
	TotalPrice float64     `json:"total_price"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	DeletedAt  *time.Time  `json:"deleted_at"`
}
