package main

import (
	"log"
	"reflect"

	"github.com/fffbbbbbb/reflact/errinfo"
)

func (e *Engine) SearchSlice(ans interface{}, s string, param ...interface{}) error {
	value := reflect.ValueOf(ans)
	modelType := value.Type()
	valueEle := value.Elem()
	if modelType.Kind() != reflect.Ptr {
		return errinfo.KindNoSruct
	}
	modelType = modelType.Elem()
	if modelType.Kind() != reflect.Slice {
		return errinfo.KindNoSlice
	}
	db := e.db
	rows, err := db.Query(s, param...)
	if err != nil {
		return err
	}

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	addSpeed := make(map[string]int)
	values := make([]interface{}, len(columns))
	nilValue := reflect.Value{}
	newEle := reflect.New(modelType.Elem()).Elem()
	for k, v := range columns {
		field := newEle.FieldByName(v)
		if field == nilValue {
			a := []byte{}
			values[k] = reflect.New(reflect.PtrTo(reflect.TypeOf(a))).Interface()
			continue
		}
		addSpeed[v] = k
		values[k] = reflect.New(reflect.PtrTo(field.Type())).Interface()
	}

	for rows.Next() {
		err = rows.Scan(values...)
		if err != nil {
			return err
		}
		for name, index := range addSpeed {
			field := newEle.FieldByName(name)
			defer func() {
				if err := recover(); err != nil {
					log.Println(err)
				}
			}()
			if !reflect.ValueOf(values[index]).Elem().IsNil() {
				field.Set(reflect.ValueOf(values[index]).Elem().Elem())
			}

		}

		valueEle.Set(reflect.Append(valueEle, newEle))

	}
	if err = rows.Err(); err != nil {
		return err
	}
	return nil
}
