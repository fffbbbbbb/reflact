package main

import (
	"errors"
	"fmt"
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

	fmt.Println("here0")

	db := e.db
	rows, err := db.Query(s, param...)
	if err != nil {
		return err
	}
	fmt.Println("here")

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return err
		}
		nilValue := reflect.Value{}
		newEle := reflect.New(modelType.Elem()).Elem()
		fmt.Println(newEle.Kind())
		for k, v := range columns {
			field := newEle.FieldByName(v)
			if field == nilValue {
				continue
			}
			defer func() {
				if err := recover(); err != nil {
					log.Println(err)
				}
			}()
			nvPtr := reflect.ValueOf(scanArgs[k])
			nv := nvPtr.Elem()
			valInterface := nv.Interface()
			fmt.Println(valInterface)
			fmt.Println(nv.Type())
			switch valInterface.(type) {
			case int, int32, int64:
				field.SetInt(valInterface.(int64))
			case bool:
				field.SetBool(valInterface.(bool))
			case float32, float64:
				field.SetFloat(valInterface.(float64))
			case string:
				field.SetString(valInterface.(string))
			case []byte:
				// field.SetString(string(valInterface.([]byte)))
				field.SetString(string(valInterface.([]byte)))
			case nil:

			default:
				return errors.New(fmt.Sprintf("no support type %v", valInterface))
			}
			// field.Set(nv.Convert(field.Type()))

		}

		valueEle.Set(reflect.Append(valueEle, newEle))

	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return nil
}
