package itempostgresql

import (
	"context"
	"fmt"
	"github.com/miladshalikar/cafe/entity"
	commonparam "github.com/miladshalikar/cafe/param/common"
	itemparam "github.com/miladshalikar/cafe/param/item"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (d *DB) GetTotalCountItemWithSearchAndFilter(
	ctx context.Context,
	search commonparam.SearchRequest,
	filter itemparam.FilterRequest,
) (uint, error) {
	const op = "itempostgresql.GetTotalCountItemWithSearchAndFilter"

	query := `SELECT COUNT(*) FROM items WHERE deleted_at IS NULL`

	var count uint
	var args []interface{}
	argIndex := 1

	if search.Search != "" {
		query += fmt.Sprintf(" AND title ILIKE $%d", argIndex)
		args = append(args, "%"+search.Search+"%")
		argIndex++
	}

	if filter.CategoryID != 0 {
		query += fmt.Sprintf(" AND category_id = $%d", argIndex)
		args = append(args, filter.CategoryID)
		argIndex++
	}

	if filter.MinPrice != 0 {
		query += fmt.Sprintf(" AND price >= $%d", argIndex)
		args = append(args, filter.MinPrice)
		argIndex++
	}

	if filter.MaxPrice != 0 {
		query += fmt.Sprintf(" AND price <= $%d", argIndex)
		args = append(args, filter.MaxPrice)
		argIndex++
	}

	if err := d.conn.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).
			WithKind(richerror.KindUnexpected)
	}
	return count, nil
}

func (d *DB) GetItemsWithPaginationAndSearchAndFilter(
	ctx context.Context,
	pagination commonparam.PaginationRequest,
	search commonparam.SearchRequest,
	filter itemparam.FilterRequest,
) ([]entity.Item, error) {
	const op = "itempostgresql.GetItemsWithPaginationAndSearchAndFilter"

	query := `SELECT * FROM items WHERE deleted_at IS NULL`

	var args []interface{}
	argIndex := 1

	if search.Search != "" {
		query += fmt.Sprintf(" AND title ILIKE $%d", argIndex)
		args = append(args, "%"+search.Search+"%")
		argIndex++
	}

	if filter.CategoryID != 0 {
		query += fmt.Sprintf(" AND category_id = $%d", argIndex)
		args = append(args, filter.CategoryID)
		argIndex++
	}

	if filter.MinPrice != 0 {
		query += fmt.Sprintf(" AND price >= $%d", argIndex)
		args = append(args, filter.MinPrice)
		argIndex++
	}

	if filter.MaxPrice != 0 {
		query += fmt.Sprintf(" AND price <= $%d", argIndex)
		args = append(args, filter.MaxPrice)
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

	var items []entity.Item

	for rows.Next() {
		item, iErr := scanItem(rows)
		if iErr != nil {
			return nil, richerror.New(op).
				WithWarpError(err).
				WithMessage(errmsg.ErrorMsgCantScanQueryResult).
				WithKind(richerror.KindUnexpected)
		}
		items = append(items, item)
	}
	if rErr := rows.Err(); rErr != nil {
		return nil, richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}

	return items, nil
}
