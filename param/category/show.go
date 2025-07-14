package categoryparam

type ShowSingleCategoryRequest struct {
	ID uint `json:"id"`
}

type ShowSingleCategoryResponse struct {
	CategoryInfo CategoryInfo `json:"category_info"`
}
