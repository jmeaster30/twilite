package twilite

import (
	"database/sql/driver"

	"github.com/jmeaster30/twilite/internal/twilib"
)

type QueryBuilder[T any] interface {
	Build() (driver.Stmt, error)
}

type SelectQueryBuilder[T any] struct {
	context *DbContext
	table   twilib.Table
}

func (s SelectQueryBuilder[T]) Build() (driver.Stmt, error) {
	return nil, nil
}
