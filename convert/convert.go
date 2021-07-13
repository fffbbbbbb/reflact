package convert

import (
	"reflect"
)

// "TEXT"
// "LONGTEXT"
// "CHAR"
// "VARCHAR"
// "MEDIUMTEXT"
// "TINYTEXT"
//"BIT"
//"BLOB"
//"DATE"
//"DATETIME"
//"DOUBLE"
//"ENUM"
//"FLOAT"
//"GEOMETRY"
//"MEDIUMINT"
//"JSON"
//"INT"
//"LONGBLOB"
//"BIGINT"
//"MEDIUMBLOB"
//"DECIMAL"
//"NULL"
//"SET"
//"SMALLINT"
//"BINARY"
//"TIME"
//"TIMESTAMP"
//"TINYINT"
//"TINYBLOB"
//"VARBINARY"
//"YEAR"

func GoTypeToDbType(goType reflect.Type, hasJson bool) string {

	switch goType.Kind() {
	case reflect.Bool:
		return "TINYINT(1)"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		return "INT"
	case reflect.Uintptr:

	case reflect.Float32, reflect.Float64:
		return "FLOAT"
	case reflect.Complex64, reflect.Complex128:

	case reflect.Array:

	case reflect.Chan:

	case reflect.Func:

	case reflect.Interface:

	case reflect.Map:

	case reflect.Ptr:

	case reflect.Slice:

	case reflect.String:
		return "VARCHAR(255)"
	case reflect.Struct:
		if hasJson {
			return "JSON"
		}
	case reflect.UnsafePointer:

	}
	return ""
}
