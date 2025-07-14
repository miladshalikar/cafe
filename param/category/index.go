package categoryparam

import "github.com/miladshalikar/cafe/param/common"

type GetCategoryRequest struct {
	Pagination common.PaginationRequest
}

type GetCategoryResponse struct {
	Pagination common.PaginationResponse
	Categories []CategoryInfo `json:"categories"`
}
