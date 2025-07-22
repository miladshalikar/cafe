package categoryparam

type AddNewCategoryRequest struct {
	Title   string `json:"title"`
	MediaID uint   `json:"media_id"`
}

type AddNewCategoryResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	MediaID uint   `json:"media_id"`
}
