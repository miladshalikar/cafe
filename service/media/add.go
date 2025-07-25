package mediaservice

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
	mediaparam "github.com/miladshalikar/cafe/param/media"
)

func (s Service) AddMedia(ctx context.Context, req mediaparam.AddMediaRequest) (mediaparam.AddMediaResponse, error) {

	media := entity.Media{
		FileName:  req.FileName,
		Size:      req.Size,
		Path:      req.Path,
		MimeType:  req.MimeType,
		IsPrivate: req.IsPrivate,
		Bucket:    req.Bucket,
	}

	media, err := s.repository.AddMedia(ctx, media)
	if err != nil {
		return mediaparam.AddMediaResponse{}, err
	}

	return mediaparam.AddMediaResponse{}, nil
}
