package entity

import "time"

type OrderItem struct {
	Id        uint       `json:"id"`
	OrderId   uint       `json:"order_id"`
	ItemId    uint       `json:"item_id"`
	Price     float64    `json:"price"`
	Quantity  uint       `json:"quantity"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
