package categoryparam

type AddNewCategoryRequest struct {
	Title string `json:"title"`
}

type AddNewCategoryResponse struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
}
