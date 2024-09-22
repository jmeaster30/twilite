package twilite

import (
	"database/sql/driver"
	"fmt"
	"reflect"
)

func getColumnType(goType reflect.Type) twiResult[string] {
	switch goType.Kind() {
	case reflect.Bool:
		return Ok("NUMERIC")
	case reflect.Float32, reflect.Float64:
		return Ok("REAL")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return Ok("INTEGER")
	case reflect.String:
		return Ok("TEXT")
	default:
		return Errorf[string]("unsupported type '%s'", goType.Kind().String())
	}
}

type twiColumnData struct {
	name        string
	columnType  string
	columnIndex uint
	goType      reflect.Type
	// TODO primaryKey    bool
	// TODO autoIncrement bool
	// TODOnullable      bool
	// TODO unique
	// TODO foreign key
}

func NewColumnData(structField reflect.StructField, index uint) twiResult[twiColumnData] {
	return OnOk(getColumnType(structField.Type).OnError(
		func(err error) twiResult[string] {
			value, ok := structField.Tag.Lookup("twiColumnType")
			return ErrorOnMissing(value, ok, fmt.Errorf("missing valid column type for field '%s'", structField.Name))
		}),
		func(columnType string) twiResult[twiColumnData] {
			return Ok(twiColumnData{
				name:        structField.Name,
				columnType:  columnType,
				columnIndex: index,
				goType:      structField.Type,
			})
		})
}

func (columnData twiColumnData) GetColumnIndex() uint {
	return columnData.columnIndex
}

func (columnData twiColumnData) GetColumnDefinition() string {
	return fmt.Sprintf("%s %s", columnData.name, columnData.columnType)
}

func (columnData twiColumnData) ToGoType(value driver.Value) twiResult[reflect.Value] {
	return Ok(reflect.ValueOf(value)) // TODO need to add nested struct types
}
