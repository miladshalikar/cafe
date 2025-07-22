package itemparam

type UpdateCategoryRequest struct {
	Id          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryId  uint    `json:"category_id"`
	MediaID     uint    `json:"media_id"`
}

type UpdateCategoryResponse struct {
	ItemInfo ItemInfo `json:"item_info"`
}
