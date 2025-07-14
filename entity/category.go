package entity

import "time"

type Category struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	MediaID   uint       `json:"media_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
