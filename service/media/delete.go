package mediaservice

import (
	"context"
	"github.com/miladshalikar/cafe/param/media"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (s Service) DeleteMedia(ctx context.Context, req mediaparam.DeleteMediaRequest) (mediaparam.DeleteMediaResponse, error) {
	const op = "mediaservice.DeleteMedia"

	media, mErr := s.repository.GetMediaByID(ctx, req.ID)
	if mErr != nil {
		return mediaparam.DeleteMediaResponse{}, richerror.New(op).WithWarpError(mErr)
	}

	cErr := s.client.Delete(ctx, media.Path)
	if cErr != nil {
		return mediaparam.DeleteMediaResponse{}, richerror.New(op).
			WithWarpError(cErr).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).
			WithKind(richerror.KindUnexpected)
	}

	rErr := s.repository.DeleteMedia(ctx, media.ID)
	if rErr != nil {
		return mediaparam.DeleteMediaResponse{}, richerror.New(op).WithWarpError(rErr)
	}

	return mediaparam.DeleteMediaResponse{}, nil
}
