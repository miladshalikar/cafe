package itemparam

type FilterRequest struct {
	CategoryID uint `json:"category_id"`
	MinPrice   uint `json:"min_price"`
	MaxPrice   uint `json:"max_price"`
}
