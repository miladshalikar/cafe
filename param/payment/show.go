package paymentparam

type ShowSinglePaymentRequest struct {
	ID uint `json:"id"`
}

type ShowSinglePaymentResponse struct {
	PaymentInfo PaymentInfo `json:"payment_info"`
}
