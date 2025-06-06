package authserviceparam

import "github.com/miladshalikar/cafe/entity"

type RegisterRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	User   entity.User `json:"user"`
	Tokens Tokens      `json:"tokens"`
}
