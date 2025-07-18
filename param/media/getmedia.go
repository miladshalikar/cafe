package mediaparam

import "github.com/miladshalikar/cafe/entity"

type GetMediaRequest struct {
	ID uint `json:"id"`
}

type GetMediaResponse struct {
	Media entity.Media `json:"media"`
}
