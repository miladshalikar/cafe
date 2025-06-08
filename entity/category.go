package entity

import "time"

type Category struct {
	Id        uint       `json:"id"`
	Title     string     `json:"title"`
	Logo      string     `json:"logo"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
