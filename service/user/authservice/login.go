package userauthservice

import (
	"context"
	"github.com/miladshalikar/cafe/param/user/authservice"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) LoginWithEmail(ctx context.Context, req userauthserviceparam.LoginWithEmailRequest) (userauthserviceparam.LoginWithEmailResponse, error) {

	user, uErr := s.repo.GetUserByEmail(ctx, req.Email)
	if uErr != nil {
		return userauthserviceparam.LoginWithEmailResponse{}, uErr
	}

	if cErr := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(req.Password)); cErr != nil {
		return userauthserviceparam.LoginWithEmailResponse{}, cErr
	}

	at, aErr := s.tokens.CreateAccessToken(user.ID)
	if aErr != nil {
		return userauthserviceparam.LoginWithEmailResponse{}, aErr
	}

	rt, rErr := s.tokens.CreateRefreshToken(user.ID)
	if rErr != nil {
		return userauthserviceparam.LoginWithEmailResponse{}, rErr
	}

	return userauthserviceparam.LoginWithEmailResponse{
		User: user,
		Tokens: userauthserviceparam.Tokens{
			AccessToken:  at,
			RefreshToken: rt,
		},
	}, nil
}
