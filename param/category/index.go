package categoryparam

import "github.com/miladshalikar/cafe/param/common"

type GetCategoriesRequest struct {
	Pagination commonparam.PaginationRequest
	Search     commonparam.SearchRequest
}

type GetCategoriesResponse struct {
	Pagination commonparam.PaginationResponse
	Categories []CategoryInfo `json:"categories"`
}
