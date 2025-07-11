package userauthservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	"github.com/miladshalikar/cafe/param/user/authservice"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Register(ctx context.Context, req userauthserviceparam.RegisterRequest) (userauthserviceparam.RegisterResponse, error) {

	u := entity.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}
	hashedPassword, hErr := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if hErr != nil {
		return userauthserviceparam.RegisterResponse{}, hErr
	}
	u.SetPassword(string(hashedPassword))

	uu, err := s.repo.CreateUser(ctx, u)
	if err != nil {
		return userauthserviceparam.RegisterResponse{}, err
	}

	at, aErr := s.tokens.CreateAccessToken(uu.ID)
	if aErr != nil {
		return userauthserviceparam.RegisterResponse{}, aErr
	}

	rt, rErr := s.tokens.CreateRefreshToken(uu.ID)
	if rErr != nil {
		return userauthserviceparam.RegisterResponse{}, rErr
	}

	return userauthserviceparam.RegisterResponse{
		User: uu,
		Tokens: userauthserviceparam.Tokens{
			AccessToken:  at,
			RefreshToken: rt,
		},
	}, nil
}
