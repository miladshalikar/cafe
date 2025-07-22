package itemparam

import commonparam "github.com/miladshalikar/cafe/param/common"

type GetItemRequest struct {
	Pagination commonparam.PaginationRequest
	Search     commonparam.SearchRequest
}

type GetItemResponse struct {
	Pagination commonparam.PaginationResponse
	Items      []ItemInfo `json:"Items"`
}
