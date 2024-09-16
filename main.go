package twilite

import "reflect"

type DbContext struct {
	databaseFile string
	tables       []Table
}

func NewDbContext(databaseFile string) DbContext {
	return DbContext{
		databaseFile: databaseFile,
		tables:       []Table{},
	}
}

func (context *DbContext) RegisterTable(reflect.Type) {

}
