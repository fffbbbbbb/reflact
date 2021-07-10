package main

import (
	"fmt"
	"testing"
)

func init() {
	username := "root"
	password := "123456"
	address := "127.0.0.1"
	dbname := "relact"
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username, password, address, dbname)
	var err error
	if testDB, err = Open(dns); err != nil {
		panic(err)
	}
}

var testDB *Engine

type Student struct{}

func (s Student) TableName() string {
	return "student_modify"
}
func TestSyncTable(t *testing.T) {
	err := testDB.SyncTable(&Student{})
	if err != nil {
		t.Error(err)
	}
}
