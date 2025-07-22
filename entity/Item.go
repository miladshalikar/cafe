package entity

import "time"

type Item struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	CategoryID  uint       `json:"category_id"`
	MediaID     *uint      `json:"media_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
