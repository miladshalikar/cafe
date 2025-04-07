package entity

import "time"

type PermissionUser struct {
	Id           uint       `json:"id"`
	UserId       uint       `json:"user_id"`
	PermissionId uint       `json:"permission_id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}
