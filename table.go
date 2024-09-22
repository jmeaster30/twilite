package twilite

import (
	"database/sql/driver"
	"reflect"

	"github.com/mattn/go-sqlite3"
)

type twiTable struct {
	name                  string
	fieldNameToColumnData map[string]twiColumnData
	fieldNameOrdering     []string
	goType                reflect.Type
}

func NewTable(structType reflect.Type) twiResult[twiTable] {
	columns := map[string]twiColumnData{}
	columnOrder := []string{}

	for i := 0; i < structType.NumField(); i++ {
		structField := structType.Field(i)
		columnDataResult := NewColumnData(structField, uint(i))
		if columnDataResult.IsError() {
			return Error[twiTable](columnDataResult.Error())
		}
		columns[structField.Name] = columnDataResult.Value()
		columnOrder = append(columnOrder, structField.Name)
	}

	return Ok(twiTable{
		name:                  structType.Name(),
		fieldNameToColumnData: columns,
		fieldNameOrdering:     columnOrder,
		goType:                structType,
	})
}

func (table twiTable) BuildTable(conn *sqlite3.SQLiteConn) twiResult[driver.Stmt] {
	query := "CREATE TABLE IF NOT EXISTS " + table.name + " ("
	for columnIndex := 0; columnIndex < len(table.fieldNameOrdering); columnIndex++ {
		columnData := table.fieldNameToColumnData[table.fieldNameOrdering[columnIndex]]
		query += columnData.GetColumnDefinition()
		if columnIndex < len(table.fieldNameOrdering)-1 {
			query += ", "
		}
	}
	query += ");"
	return ToResult(conn.Prepare(query))
}

func (table *twiTable) ToGoType(row []driver.Value) twiResult[reflect.Value] {
	result := reflect.New(table.goType)
	for fieldIndex := 0; fieldIndex < table.goType.NumField(); fieldIndex++ {
		field := table.goType.Field(fieldIndex)
		columnData := table.fieldNameToColumnData[field.Name]
		columnValue := columnData.ToGoType(row[columnData.GetColumnIndex()])
		if columnValue.IsError() {
			return Error[reflect.Value](columnValue.Error())
		}
		result.FieldByName(field.Name).Set(columnValue.Value())
	}
	return Ok(result)
}
