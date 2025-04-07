package entity

import "time"

type User struct {
	Id          uint   `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	password    string
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) SetPassword(p string) {
	u.password = p
}
