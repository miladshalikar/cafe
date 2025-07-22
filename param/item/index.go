package itemparam

import commonparam "github.com/miladshalikar/cafe/param/common"

type GetItemsRequest struct {
	Pagination commonparam.PaginationRequest
	Search     commonparam.SearchRequest
}

type GetItemsResponse struct {
	Pagination commonparam.PaginationResponse
	Items      []ItemInfo `json:"Items"`
}
