package paymentparam

type SimulatePaymentResultRequest struct {
	PaymentID uint `json:"payment_id"`
	Success   bool `json:"success"`
}

type SimulatePaymentResultResponse struct {
}
