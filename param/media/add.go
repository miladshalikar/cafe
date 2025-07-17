package mediaparam

type AddMediaRequest struct {
	FileName  string `json:"file_name"`
	Size      uint   `json:"size"`
	Path      string `json:"path"`
	MimeType  string `json:"mime_type"`
	IsPrivate bool   `json:"is_private"`
	Bucket    string `json:"bucket"`
}

type AddMediaResponse struct {
}
