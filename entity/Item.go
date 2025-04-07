package entity

import "time"

type Item struct {
	Id          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	CategoryId  uint       `json:"category_id"`
	Logo        string     `json:"logo"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
