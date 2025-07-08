package categoryparam

type UpdateCategoryRequest struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Logo  string `json:"logo"`
}

type UpdateCategoryResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Logo  string `json:"logo"`
}
