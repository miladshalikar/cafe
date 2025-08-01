package userprofileservice

import (
	"context"
	"github.com/miladshalikar/cafe/param/user/profile"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (s Service) GetUserByID(ctx context.Context, req userprofileserviceparam.UserProfileRequest) (userprofileserviceparam.UserProfileResponse, error) {
	const op = "userprofileservice.GetUserByID"

	user, err := s.repo.GetUserByID(ctx, req.Id)
	if err != nil {
		return userprofileserviceparam.UserProfileResponse{}, richerror.New(op).WithWarpError(err)
	}

	return userprofileserviceparam.UserProfileResponse{
		User: user,
	}, nil
}
