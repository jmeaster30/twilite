package twilite

import "reflect"

type Table struct {
	selectStmt string
	insertStmt string
	name       string
	columns    string
	goType     reflect.Type
}
