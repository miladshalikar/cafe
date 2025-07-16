package categoryservice

import (
	"context"
	categoryparam "github.com/miladshalikar/cafe/param/category"
	"github.com/miladshalikar/cafe/param/common"
)

func (s Service) GetCategories(ctx context.Context, req categoryparam.GetCategoryRequest) (categoryparam.GetCategoryResponse, error) {

	total, tErr := s.repo.GetTotalCountCategory(ctx, req.Search)
	if tErr != nil {
		return categoryparam.GetCategoryResponse{}, tErr
	}

	categories, cErr := s.repo.GetCategoriesWithPagination(ctx, req.Pagination.GetPageSize(), req.Pagination.GetOffset(), req.Search)
	if cErr != nil {
		return categoryparam.GetCategoryResponse{}, cErr
	}

	var categoriesInfo []categoryparam.CategoryInfo
	for _, category := range categories {

		categoriesInfo = append(categoriesInfo, categoryparam.CategoryInfo{
			ID:      category.ID,
			Title:   category.Title,
			MediaID: category.MediaID,
		})

	}

	return categoryparam.GetCategoryResponse{
		Pagination: commonparam.PaginationResponse{
			PageSize:   req.Pagination.GetPageSize(),
			PageNumber: req.Pagination.GetPageNumber(),
			Total:      total,
		},
		Categories: categoriesInfo,
	}, nil

}
