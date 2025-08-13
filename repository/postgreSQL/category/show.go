package categorypostgresql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/miladshalikar/cafe/entity"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (d *DB) GetCategoryByID(ctx context.Context, id uint) (entity.Category, error) {
	const op = "categorypostgresql.GetCategoryByID"

	query := `SELECT * FROM categories WHERE id = $1 AND deleted_at IS NULL`

	row := d.conn.QueryRowContext(ctx, query, id)

	category, err := scanCategory(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Category{}, richerror.New(op).
				WithWarpError(err).
				WithMessage(errmsg.ErrorMsgNotFound).
				WithKind(richerror.KindNotFound)
		}
		return entity.Category{}, richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}

	return category, nil

}

func (d *DB) CheckCategoryIsExistByID(ctx context.Context, id uint) (bool, error) {
	const op = "categorypostgresql.CheckCategoryIsExistByID"

	query := `SELECT * FROM categories WHERE id = $1 AND deleted_at IS NULL`

	row := d.conn.QueryRowContext(ctx, query, id)

	_, err := scanCategory(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected).
			WithMeta(map[string]interface{}{"id": id})
	}

	return true, nil
}

func (d *DB) CheckCategoryIsExistByTitle(ctx context.Context, title string) (bool, error) {
	const op = "categorypostgresql.CheckCategoryIsExistByTitle"

	query := `SELECT * FROM categories WHERE title = $1 AND deleted_at IS NULL`

	row := d.conn.QueryRowContext(ctx, query, title)

	_, err := scanCategory(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected).
			WithMeta(map[string]interface{}{"title": title})
	}

	return true, nil
}
