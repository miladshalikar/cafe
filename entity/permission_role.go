package entity

import "time"

type PermissionRole struct {
	ID           uint       `json:"id"`
	RoleId       uint       `json:"role_id"`
	PermissionId uint       `json:"permission_id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}
