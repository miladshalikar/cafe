package mediaparam

type GetURLRequest struct {
	ID uint `json:"id"`
}

type GetURLResponse struct {
	URL      string `json:"url"`
	MimeType string `json:"mime_type"`
}
