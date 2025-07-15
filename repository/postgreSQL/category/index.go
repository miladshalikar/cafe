package categorypostgresql

import (
	"context"
	"github.com/miladshalikar/cafe/entity"
)

func (d *DB) GetTotalCountArea(ctx context.Context) (uint, error) {
	panic("aa")
}

func (d *DB) GetAreasWithPagination(ctx context.Context, pageSize, offset uint) ([]entity.Category, error) {
	panic("aaa")
}
