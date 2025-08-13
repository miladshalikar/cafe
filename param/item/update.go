package itemparam

type UpdateItemRequest struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    uint    `json:"quantity"`
	CategoryID  uint    `json:"category_id"`
	MediaID     uint    `json:"media_id"`
}

type UpdateItemResponse struct {
	ItemInfo ItemInfo `json:"item_info"`
}
