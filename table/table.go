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

//传入参数为Struct
func GetTableName(opt interface{}, nameFunc func(a string) string) string {
	v := reflect.ValueOf(opt)
	modelType := v.Type()
	if modelType.Kind() != reflect.Struct {
		return ""
	}
	tableName := ""
	m, ok := modelType.MethodByName("TableName")
	if ok {
		ans := m.Func.Call([]reflect.Value{v})
		if ans[0].Kind() == reflect.String {
			tableName = ans[0].String()
		}
	} else if nameFunc != nil {
		tableName = nameFunc(modelType.Name())
	}
	return tableName
}
