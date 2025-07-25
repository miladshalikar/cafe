package mediaservice

import (
	"context"
	"github.com/gabriel-vasile/mimetype"
	"github.com/miladshalikar/cafe/entity"
	mediaparam "github.com/miladshalikar/cafe/param/media"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func (s Service) UploadMedia(ctx context.Context, req mediaparam.UploadMediaRequest) (mediaparam.UploadMediaResponse, error) {

	fileName := strconv.FormatInt(time.Now().UnixNano(), 10) + filepath.Ext(req.Filename)

	file, fErr := req.Open()
	if fErr != nil {
		return mediaparam.UploadMediaResponse{}, fErr
	}
	defer file.Close()

	mType, dErr := mimetype.DetectReader(file)
	if dErr != nil {
		return mediaparam.UploadMediaResponse{}, dErr
	}

	datePath := strings.ReplaceAll(time.Now().UTC().Format(time.DateOnly), "-", "/")
	fullPath := "public/" + datePath + "/" + fileName

	media := entity.Media{
		FileName:  fileName,
		Size:      uint(req.Size),
		Path:      fullPath,
		MimeType:  mType.String(),
		IsPrivate: req.IsPrivate,
		Bucket:    req.Bucket,
	}

	uErr := s.client.Upload(ctx, *req.FileHeader, fullPath)
	if uErr != nil {
		return mediaparam.UploadMediaResponse{}, uErr
	}

	insertedMedia, iErr := s.repository.AddMedia(ctx, media)
	if iErr != nil {
		return mediaparam.UploadMediaResponse{}, uErr
	}

	res, gErr := s.GetURLMedia(ctx, mediaparam.GetURLRequest{ID: insertedMedia.ID})
	if gErr != nil {
		return mediaparam.UploadMediaResponse{}, gErr
	}

	return mediaparam.UploadMediaResponse{
		ID:       insertedMedia.ID,
		URL:      res.URL,
		MimeType: res.MimeType,
	}, nil

}
