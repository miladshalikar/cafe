package mediaservice

import (
	"context"
	mediaparam "github.com/miladshalikar/cafe/param/media"
)

func (s Service) GetURLMedia(ctx context.Context, req mediaparam.GetURLRequest) (mediaparam.GetURLResponse, error) {

	media, mErr := s.repository.GetMediaByID(ctx, req.ID)
	if mErr != nil {
		return mediaparam.GetURLResponse{}, mErr
	}

	url, err := s.client.GetURL(ctx, media.Path)
	if err != nil {
		return mediaparam.GetURLResponse{}, err
	}

	return mediaparam.GetURLResponse{
		URL:      url,
		MimeType: media.MimeType,
	}, nil
}
