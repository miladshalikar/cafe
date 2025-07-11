package entity

import "time"

type RoleUser struct {
	ID        uint       `json:"id"`
	RoleId    uint       `json:"role_id"`
	UserId    uint       `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
