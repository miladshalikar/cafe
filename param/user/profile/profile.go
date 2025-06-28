package userprofileserviceparam

import "github.com/miladshalikar/cafe/entity"

type UserProfileRequest struct {
	Id int `json:"id"`
}

type UserProfileResponse struct {
	User entity.User `json:"user"`
}
