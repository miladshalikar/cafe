package categorypostgresql

import (
	"context"
	"fmt"
	"github.com/miladshalikar/cafe/entity"
	commonparam "github.com/miladshalikar/cafe/param/common"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (d *DB) GetTotalCountCategoryWithSearch(ctx context.Context, search commonparam.SearchRequest) (uint, error) {
	const op = "categorypostgresql.GetTotalCountCategoryWithSearch"

	query := `SELECT COUNT(*) FROM categories WHERE deleted_at IS NULL`

	var count uint
	var args []interface{}

	if search.Search != "" {
		query += " AND title ILIKE $1"
		args = append(args, "%"+search.Search+"%")
	}

	if err := d.conn.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).
			WithKind(richerror.KindUnexpected)
	}
	return count, nil

}

func (d *DB) GetCategoriesWithPaginationAndSearch(ctx context.Context, pagination commonparam.PaginationRequest, search commonparam.SearchRequest) ([]entity.Category, error) {
	const op = "categorypostgresql.GetCategoriesWithPaginationAndSearch"

	query := `SELECT * FROM categories WHERE deleted_at IS NULL`
	var args []interface{}
	argIndex := 1

	if search.Search != "" {
		query += fmt.Sprintf(" AND title ILIKE $%d", argIndex)
		args = append(args, "%"+search.Search+"%")
		argIndex++
	}

	query += " ORDER BY id ASC"
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, pagination.GetPageSize(), pagination.GetOffset())

	rows, err := d.conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).
			WithKind(richerror.KindUnexpected)
	}
	defer rows.Close()

	var categories []entity.Category

	for rows.Next() {
		category, cErr := scanCategory(rows)
		if cErr != nil {
			return nil, richerror.New(op).
				WithWarpError(err).
				WithMessage(errmsg.ErrorMsgCantScanQueryResult).
				WithKind(richerror.KindUnexpected)
		}
		categories = append(categories, category)
	}
	if rErr := rows.Err(); rErr != nil {
		return nil, richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}

	return categories, nil
}
