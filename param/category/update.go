package categoryparam

type UpdateCategoryRequest struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	MediaID uint   `json:"media_id"`
}

type UpdateCategoryResponse struct {
	Category CategoryInfo `json:"category"`
}
