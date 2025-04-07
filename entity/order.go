package entity

import "time"

type Order struct {
	Id         uint       `json:"id"`
	UserId     uint       `json:"user_id"`
	Status     string     `json:"status"`
	TotalPrice float64    `json:"total_price"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
