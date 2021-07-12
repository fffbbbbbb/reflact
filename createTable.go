package main

import (
	"database/sql"
	"reflect"
	"strings"

	"github.com/fffbbbbbb/reflact/errinfo"
)

func (db *Engine) SyncTable(opt ...interface{}) error {
	for _, v := range opt {
		err := db.syncTable(v, db.nameFunc)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Engine) syncTable(opt interface{}, nameFunc func(a string) string) error {
	db := e.db
	v := reflect.ValueOf(opt)
	modelType := v.Type()
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}
	if modelType.Kind() != reflect.Struct {
		return errinfo.KindNoSruct
	}
	tableName := ""
	m, ok := modelType.MethodByName("TableName")
	if ok {
		ans := m.Func.Call([]reflect.Value{v.Elem()})
		if ans[0].Kind() == reflect.String {
			tableName = ans[0].String()
		}
	} else {
		tableName = nameFunc(modelType.Name())
	}
	exist, err := hasTable(db, tableName)
	if err != nil {
		return nil
	}
	if exist {
		return nil
	}
	return nil
}

func hasTable(db *sql.DB, tableName string) (bool, error) {
	s := "select count(*) from information_schema.TABLES where TABLE_NAME=? "
	count := 0
	err := db.QueryRow(s, tableName).Scan(&count)
	if err != nil {
		return false, err
	}
	if count != 0 {
		return true, nil
	}
	return false, nil
}

func nameFunc(name string) string {
	s := []string{}
	now := []byte{}
	for i := 0; i < len(name); i++ {
		if isUpper(name[i]) && len(now) != 0 {
			nowStr := strings.ToLower(string(now))
			s = append(s, string(nowStr))
			now = []byte{}

		}
		now = append(now, name[i])
	}
	nowStr := strings.ToLower(string(now))
	s = append(s, string(nowStr))
	// fmt.Println(s)
	return strings.Join(s, "_")
}

func isUpper(a byte) bool {
	if a >= 'A' && a <= 'Z' {
		return true
	}
	return false
}
