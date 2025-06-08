package userauthservice

import (
	"context"
	param "github.com/miladshalikar/cafe/param/authservice"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) LoginWithEmail(ctx context.Context, req param.LoginWithEmailRequest) (param.LoginWithEmailResponse, error) {

	user, uErr := s.Repo.GetUserByEmail(ctx, req.Email)
	if uErr != nil {
		return param.LoginWithEmailResponse{}, uErr
	}

	if cErr := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(req.Password)); cErr != nil {
		return param.LoginWithEmailResponse{}, cErr
	}

	at, aErr := s.Tokens.CreateAccessToken(user.Id)
	if aErr != nil {
		return param.LoginWithEmailResponse{}, aErr
	}

	rt, rErr := s.Tokens.CreateRefreshToken(user.Id)
	if rErr != nil {
		return param.LoginWithEmailResponse{}, rErr
	}

	return param.LoginWithEmailResponse{
		User: user,
		Tokens: param.Tokens{
			AccessToken:  at,
			RefreshToken: rt,
		},
	}, nil
}
