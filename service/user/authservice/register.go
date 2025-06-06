package userauthservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	param "github.com/miladshalikar/cafe/param/authservice"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Register(ctx context.Context, req param.RegisterRequest) (param.RegisterResponse, error) {

	u := entity.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}
	hashedPassword, hErr := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if hErr != nil {
		return param.RegisterResponse{}, hErr
	}
	u.SetPassword(string(hashedPassword))

	uu, err := s.Repo.CreateUser(ctx, u)
	if err != nil {
		return param.RegisterResponse{}, err
	}

	at, aErr := s.Tokens.CreateAccessToken(uu.Id)
	if aErr != nil {
		return param.RegisterResponse{}, aErr
	}

	rt, rErr := s.Tokens.CreateRefreshToken(uu.Id)
	if rErr != nil {
		return param.RegisterResponse{}, rErr
	}

	return param.RegisterResponse{
		UserID: uu.Id,
		Tokens: param.Tokens{
			AccessToken:  at,
			RefreshToken: rt,
		},
	}, nil
}
