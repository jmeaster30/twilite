package twilite

import (
	"reflect"

	"twilite/twilib"
)

type DbContext struct {
	databaseFile               string
	tables                     map[string]twilib.Table
	createIdColumnIfNotPresent bool
}

func NewDbContext(databaseFile string) DbContext {
	return DbContext{
		databaseFile:               databaseFile,
		tables:                     map[string]twilib.Table{},
		createIdColumnIfNotPresent: false,
	}
}

func (context *DbContext) CreateIdColumnIfNotPresent() {
	context.createIdColumnIfNotPresent = true
}

func (context *DbContext) InitializeTables() error {
	return nil
}

func RegisterTable[T any](context *DbContext) error {
	var zero [0]T
	structType := reflect.TypeOf(zero).Elem()

	tableResult := twilib.NewTable(structType)
	if tableResult.IsError() {
		return tableResult.Error()
	}
	context.tables[structType.Name()] = tableResult.Value()
	return nil
}

func SelectQuery[T any](context *DbContext) QueryBuilder[T] {
	var zero [0]T
	structType := reflect.TypeOf(zero).Elem()

	return SelectQueryBuilder[T]{
		context: context,
		table:   context.tables[structType.Name()], // TODO: Add error for tables that don't exist
	}
}
