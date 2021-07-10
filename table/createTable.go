package table

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
)

var (
	KindNoSruct = errors.New("param is not struct")
)

func SyncTable(db *sql.DB, opt interface{}) error {
	v := reflect.ValueOf(opt)
	modelType := v.Type()
	fmt.Println(modelType)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}
	fmt.Println(modelType.Kind())
	if modelType.Kind() != reflect.Struct {
		return KindNoSruct
	}
	tableName := "testTableName"
	m, ok := modelType.MethodByName("TableName")
	if ok {
		ans := m.Func.Call([]reflect.Value{v.Elem()})
		if ans[0].Kind() == reflect.String {
			tableName = ans[0].String()
		}
	}
	fmt.Println(tableName)
	return nil
}

func hasTable(db *sql.DB, tableName string) {

}
