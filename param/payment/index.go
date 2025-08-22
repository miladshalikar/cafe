package paymentparam

import commonparam "github.com/miladshalikar/cafe/param/common"

type GetPaymentsByOrderIDRequest struct {
	Pagination commonparam.PaginationRequest
	OrderID    uint `json:"order_id"`
}

type GetPaymentsByOrderIDResponse struct {
	Pagination commonparam.PaginationResponse
	Payments   []PaymentInfo `json:"payment_info"`
}
