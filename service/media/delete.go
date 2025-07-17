package mediaservice

import (
	"context"
	"github.com/miladshalikar/cafe/param/media"
)

func (s Service) DeleteMedia(ctx context.Context, req mediaparam.DeleteMediaRequest) (mediaparam.DeleteMediaResponse, error) {

	media, mErr := s.repository.GetMediaByID(ctx, req.ID)
	if mErr != nil {
		return mediaparam.DeleteMediaResponse{}, mErr
	}

	cErr := s.client.Delete(ctx, media.Path)
	if cErr != nil {
		return mediaparam.DeleteMediaResponse{}, cErr
	}

	rErr := s.repository.DeleteMedia(ctx, media.ID)
	if rErr != nil {
		return mediaparam.DeleteMediaResponse{}, rErr
	}

	return mediaparam.DeleteMediaResponse{}, nil
}
