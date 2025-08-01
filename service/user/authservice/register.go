package userauthservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	"github.com/miladshalikar/cafe/param/user/authservice"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Register(ctx context.Context, req userauthserviceparam.RegisterRequest) (userauthserviceparam.RegisterResponse, error) {
	const op = "userauthservice.Register"

	u := entity.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}
	hashedPassword, hErr := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if hErr != nil {
		return userauthserviceparam.RegisterResponse{},
			richerror.New(op).
				WithWarpError(hErr).
				WithMessage(errmsg.ErrorMsgSomethingWentWrong).
				WithKind(richerror.KindUnexpected)
	}
	u.SetPassword(string(hashedPassword))

	uu, err := s.repo.CreateUser(ctx, u)
	if err != nil {
		return userauthserviceparam.RegisterResponse{}, richerror.New(op).WithWarpError(err)
	}

	at, aErr := s.tokens.CreateAccessToken(uu.ID)
	if aErr != nil {
		return userauthserviceparam.RegisterResponse{}, richerror.New(op).WithWarpError(aErr)
	}

	rt, rErr := s.tokens.CreateRefreshToken(uu.ID)
	if rErr != nil {
		return userauthserviceparam.RegisterResponse{}, richerror.New(op).WithWarpError(rErr)
	}

	return userauthserviceparam.RegisterResponse{
		User: uu,
		Tokens: userauthserviceparam.Tokens{
			AccessToken:  at,
			RefreshToken: rt,
		},
	}, nil
}
