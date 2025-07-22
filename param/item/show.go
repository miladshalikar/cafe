package itemparam

type ShowSingleItemRequest struct {
	ID uint `json:"id"`
}

type ShowSingleItemResponse struct {
	ItemInfo ItemInfo `json:"item_info"`
}
