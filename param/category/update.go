package categoryparam

type UpdateCategoryRequest struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type UpdateCategoryResponse struct {
	Category CategoryInfo `json:"category"`
}
