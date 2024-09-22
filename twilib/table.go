package twilib

import (
	"database/sql/driver"
	"reflect"
	"twilite/twiutil"

	"github.com/mattn/go-sqlite3"
)

type Table struct {
	name                  string
	fieldNameToColumnData map[string]ColumnData
	fieldNameOrdering     []string
	goType                reflect.Type
}

func NewTable(structType reflect.Type) twiutil.Result[Table] {
	columns := map[string]ColumnData{}
	columnOrder := []string{}

	for i := 0; i < structType.NumField(); i++ {
		structField := structType.Field(i)
		columnDataResult := NewColumnData(structField, uint(i))
		if columnDataResult.IsError() {
			return twiutil.Error[Table](columnDataResult.Error())
		}
		columns[structField.Name] = columnDataResult.Value()
		columnOrder = append(columnOrder, structField.Name)
	}

	return twiutil.Ok(Table{
		name:                  structType.Name(),
		fieldNameToColumnData: columns,
		fieldNameOrdering:     columnOrder,
		goType:                structType,
	})
}

func (table Table) BuildTable(conn *sqlite3.SQLiteConn) twiutil.Result[driver.Stmt] {
	query := "CREATE TABLE IF NOT EXISTS " + table.name + " ("
	for columnIndex := 0; columnIndex < len(table.fieldNameOrdering); columnIndex++ {
		columnData := table.fieldNameToColumnData[table.fieldNameOrdering[columnIndex]]
		query += columnData.GetColumnDefinition()
		if columnIndex < len(table.fieldNameOrdering)-1 {
			query += ", "
		}
	}
	query += ");"
	return twiutil.ToResult(conn.Prepare(query))
}

func (table *Table) ToGoType(row []driver.Value) twiutil.Result[reflect.Value] {
	result := reflect.New(table.goType)
	for fieldIndex := 0; fieldIndex < table.goType.NumField(); fieldIndex++ {
		field := table.goType.Field(fieldIndex)
		columnData := table.fieldNameToColumnData[field.Name]
		columnValue := columnData.ToGoType(row[columnData.GetColumnIndex()])
		if columnValue.IsError() {
			return twiutil.Error[reflect.Value](columnValue.Error())
		}
		result.FieldByName(field.Name).Set(columnValue.Value())
	}
	return twiutil.Ok(result)
}
