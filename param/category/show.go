package categoryparam

type ShowSingleCategoryRequest struct {
	ID uint `json:"id"`
}

type ShowSingleCategoryResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Logo  string `json:"logo"`
}
