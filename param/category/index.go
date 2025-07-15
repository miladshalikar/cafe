package categoryparam

import "github.com/miladshalikar/cafe/param/common"

type GetCategoryRequest struct {
	Pagination commonparam.PaginationRequest
}

type GetCategoryResponse struct {
	Pagination commonparam.PaginationResponse
	Categories []CategoryInfo `json:"categories"`
}
