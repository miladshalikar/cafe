package categoryparam

type AddNewCategoryRequest struct {
	Title string `json:"title"`
	Logo  string `json:"logo"`
}

type AddNewCategoryResponse struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
	Logo  string `json:"logo"`
}
