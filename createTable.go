package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/fffbbbbbb/reflact/convert"
	"github.com/fffbbbbbb/reflact/errinfo"
	"github.com/fffbbbbbb/reflact/table"
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
		return errinfo.KindNoStruct
	}
	tableName := table.GetTableName(v.Elem().Interface(), nameFunc)
	exist, err := hasTable(db, tableName)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}
	tableInfo, err := TableDescription(opt, tableName, e.hasJson)
	if err != nil {
		return err
	}
	createTableSql := tableInfo.MakeCreateSQL()
	_, err = db.Exec(createTableSql)
	if err != nil {
		return err
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

func TableDescription(opt interface{}, tableName string, hasJson bool) (*table.Table, error) {
	v := reflect.ValueOf(opt)
	modelType := v.Type()
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}
	if modelType.Kind() != reflect.Struct {
		return nil, errinfo.KindNoStruct
	}
	t := &table.Table{
		TableName: tableName,
		Field:     make([]table.Field, modelType.NumField()),
	}

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		t.Field[i].Name = field.Name
		t.Field[i].GoType = field.Type
		t.Field[i].DbType = convert.GoTypeToDbType(t.Field[i].GoType, hasJson)
		if t.Field[i].DbType == "" {
			return t, fmt.Errorf("unsupport GO type(%v) to create table ", t.Field[i].GoType)
		}
		t.Field[i].Constraint = field.Tag.Get("form")
	}

	return t, nil
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
	return strings.Join(s, "_")
}

func isUpper(a byte) bool {
	if a >= 'A' && a <= 'Z' {
		return true
	}
	return false
}
