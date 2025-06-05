package entity

import "time"

type User struct {
	Id          uint
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	password    string
	isSuperUser bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) SetPassword(p string) {
	u.password = p
}
