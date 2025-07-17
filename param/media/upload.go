package mediaparam

import "mime/multipart"

type UploadMediaRequest struct {
	*multipart.FileHeader `json:"file"`
	IsPrivate             bool
	Bucket                string
}

type UploadMediaResponse struct {
	ID       uint   `json:"id"`
	URL      string `json:"url"`
	MimeType string `json:"mime_type"`
}
