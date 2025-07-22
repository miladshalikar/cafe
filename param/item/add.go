package itemparam

type AddNewItemRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryId  uint    `json:"category_id"`
	MediaID     uint    `json:"media_id"`
}

type AddNewItemResponse struct {
	Id          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryId  uint    `json:"category_id"`
	MediaID     uint    `json:"media_id"`
}
