package authserviceparam

import "github.com/miladshalikar/cafe/entity"

type LoginWithEmailRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginWithEmailResponse struct {
	User   entity.User `json:"user"`
	Tokens Tokens      `json:"tokens"`
}
