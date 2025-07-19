package entity

import "time"

var ValidExt = []string{"mp4", "png", "jpg", "jpeg"}

const (
	MaxFileUploadSize          = 5 * 1024 * 1024
	FileLinkExpirationDuration = 60 * int64(time.Second) * 60 * 7
)

type Media struct {
	ID        uint       `json:"id"`
	FileName  string     `json:"file_name"`
	Size      uint       `json:"size"`
	Path      string     `json:"path"`
	MimeType  string     `json:"mime_type"`
	IsPrivate bool       `json:"is_private"`
	Bucket    string     `json:"bucket"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
