package userauthservice

import (
	"context"
	"github.com/miladshalikar/cafe/param/user/authservice"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) LoginWithEmail(ctx context.Context, req userauthserviceparam.LoginWithEmailRequest) (userauthserviceparam.LoginWithEmailResponse, error) {
	const op = "userauthservice.LoginWithEmail"

	user, uErr := s.repo.GetUserByEmail(ctx, req.Email)
	if uErr != nil {
		return userauthserviceparam.LoginWithEmailResponse{}, richerror.New(op).WithWarpError(uErr)
	}

	if cErr := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(req.Password)); cErr != nil {
		return userauthserviceparam.LoginWithEmailResponse{}, richerror.New(op).
			WithWarpError(cErr).
			WithMessage(errmsg.ErrorMsgEmailOrPassIsIncorrect).
			WithKind(richerror.KindForbidden)
	}

	at, aErr := s.tokens.CreateAccessToken(user.ID)
	if aErr != nil {
		return userauthserviceparam.LoginWithEmailResponse{}, richerror.New(op).WithWarpError(aErr).WithKind(richerror.KindUnexpected)
	}

	rt, rErr := s.tokens.CreateRefreshToken(user.ID)
	if rErr != nil {
		return userauthserviceparam.LoginWithEmailResponse{}, richerror.New(op).WithWarpError(rErr).WithKind(richerror.KindUnexpected)
	}

	return userauthserviceparam.LoginWithEmailResponse{
		User: user,
		Tokens: userauthserviceparam.Tokens{
			AccessToken:  at,
			RefreshToken: rt,
		},
	}, nil
}
