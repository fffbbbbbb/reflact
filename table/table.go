package table

import (
	"fmt"
	"reflect"
	"strings"
)

type TableDescInfo struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

type Table struct {
	TableName string
	Field     []Field
}

type Field struct {
	Name       string
	GoType     reflect.Type
	DbType     string
	Constraint string
}

type Version struct {
	Version string
}

func (t *Table) MakeCreateSQL() string {
	templ := `create table %s (%s)`
	fieldSilce := make([]string, len(t.Field))
	for k, v := range t.Field {
		fieldSilce[k] = v.Name + " " + v.DbType + " " + " " + v.Constraint
	}
	return fmt.Sprintf(templ, t.TableName, strings.Join(fieldSilce, ",\n"))
}
