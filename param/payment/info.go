package paymentparam

import "github.com/miladshalikar/cafe/entity"

type PaymentInfo struct {
	ID      uint                 `json:"id"`
	OrderID uint                 `json:"order_id"`
	Amount  float64              `json:"amount"`
	Status  entity.PaymentStatus `json:"status"`
	Method  string               `json:"method"`
}
