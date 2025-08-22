package paymentparam

import "github.com/miladshalikar/cafe/entity"

type RefundPaymentRequest struct {
	PaymentID uint `json:"payment_id"`
}

type RefundPaymentResponse struct {
	PaymentID uint                 `json:"payment_id"`
	Status    entity.PaymentStatus `json:"status"`
	Message   string               `json:"message"`
}
