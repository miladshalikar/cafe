package userprofileservice

import (
	"context"
	"github.com/miladshalikar/cafe/param/user/profile"
)

func (s Service) GetUserByID(ctx context.Context, req userprofileserviceparam.UserProfileRequest) (userprofileserviceparam.UserProfileResponse, error) {

	user, err := s.repo.GetUserByID(ctx, req.Id)
	if err != nil {
		return userprofileserviceparam.UserProfileResponse{}, err
	}

	return userprofileserviceparam.UserProfileResponse{
		User: user,
	}, nil
}
