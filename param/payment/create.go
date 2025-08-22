package paymentparam

import "github.com/miladshalikar/cafe/entity"

type CreatePaymentRequest struct {
	OrderID uint    `json:"order_id"`
	Amount  float64 `json:"amount"`
	Method  string  `json:"method"`
}

type CreatePaymentResponse struct {
	ID      uint                 `json:"id"`
	OrderID uint                 `json:"order_id"`
	Amount  float64              `json:"amount"`
	Status  entity.PaymentStatus `json:"status"`
	Method  string               `json:"method"`
}
