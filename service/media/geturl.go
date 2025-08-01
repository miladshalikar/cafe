package mediaservice

import (
	"context"
	mediaparam "github.com/miladshalikar/cafe/param/media"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (s Service) GetURLMedia(ctx context.Context, req mediaparam.GetURLRequest) (mediaparam.GetURLResponse, error) {
	const op = "mediaservice.GetURLMedia"

	media, mErr := s.repository.GetMediaByID(ctx, req.ID)
	if mErr != nil {
		return mediaparam.GetURLResponse{}, richerror.New(op).WithWarpError(mErr)
	}

	url, err := s.client.GetURL(ctx, media.Path)
	if err != nil {
		return mediaparam.GetURLResponse{}, richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).
			WithKind(richerror.KindUnexpected)
	}

	return mediaparam.GetURLResponse{
		URL:      url,
		MimeType: media.MimeType,
	}, nil
}
