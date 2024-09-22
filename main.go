package twilite

import (
	"reflect"

	"twilite/twilib"
)

type DbContext struct {
	databaseFile               string
	tables                     []twilib.Table
	createIdColumnIfNotPresent bool
}

func NewDbContext(databaseFile string) DbContext {
	return DbContext{
		databaseFile:               databaseFile,
		tables:                     []twilib.Table{},
		createIdColumnIfNotPresent: false,
	}
}

func (context *DbContext) CreateIdColumnIfNotPresent() {
	context.createIdColumnIfNotPresent = true
}

func (context *DbContext) InitializeTables() error {
	return nil
}

func (context *DbContext) RegisterTable(structType reflect.Type) error {
	tableResult := twilib.NewTable(structType)
	if tableResult.IsError() {
		return tableResult.Error()
	}
	context.tables = append(context.tables, tableResult.Value())
	return nil
}
