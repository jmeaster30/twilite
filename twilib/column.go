package twilib

import (
	"database/sql/driver"
	"fmt"
	"reflect"

	"github.com/jmeaster30/twilite/twiutil"
)

func getColumnType(goType reflect.Type) twiutil.Result[string] {
	switch goType.Kind() {
	case reflect.Bool:
		return twiutil.Ok("NUMERIC")
	case reflect.Float32, reflect.Float64:
		return twiutil.Ok("REAL")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return twiutil.Ok("INTEGER")
	case reflect.String:
		return twiutil.Ok("TEXT")
	default:
		return twiutil.Errorf[string]("unsupported type '%s'", goType.Kind().String())
	}
}

type ColumnData struct {
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

func NewColumnData(structField reflect.StructField, index uint) twiutil.Result[ColumnData] {
	return twiutil.OnOk(getColumnType(structField.Type).OnError(
		func(err error) twiutil.Result[string] {
			value, ok := structField.Tag.Lookup("twiColumnType")
			return twiutil.ErrorOnMissing(value, ok, fmt.Errorf("missing valid column type for field '%s'", structField.Name))
		}),
		func(columnType string) twiutil.Result[ColumnData] {
			return twiutil.Ok(ColumnData{
				name:        structField.Name,
				columnType:  columnType,
				columnIndex: index,
				goType:      structField.Type,
			})
		})
}

func (columnData ColumnData) GetColumnIndex() uint {
	return columnData.columnIndex
}

func (columnData ColumnData) GetColumnDefinition() string {
	return fmt.Sprintf("%s %s", columnData.name, columnData.columnType)
}

func (columnData ColumnData) ToGoType(value driver.Value) twiutil.Result[reflect.Value] {
	return twiutil.Ok(reflect.ValueOf(value)) // TODO need to add nested struct types
}
