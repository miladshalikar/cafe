package mediaparam

import "mime/multipart"

type UploadMediaRequest struct {
	*multipart.FileHeader `json:"file"`
	IsPrivate             bool   `json:"is_private"`
	Bucket                string `json:"bucket"`
}

type UploadMediaResponse struct {
	ID       uint   `json:"id"`
	URL      string `json:"url"`
	MimeType string `json:"mime_type"`
}
