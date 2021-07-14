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

type StudentTest struct {
	ID      int64
	Book    string
	Publish string
	Price   float64
}

func (s StudentTest) TableName() string {
	return "book"
}
func TestSyncTable(t *testing.T) {
	err := testDB.SyncTable(&StudentTest{})
	if err != nil {
		t.Error(err)
	}
}

func TestFindStudent(t *testing.T) {
	ans := &[]StudentTest{}
	if err := testDB.Where("Publish='renmin'").Find(ans); err != nil {
		t.Error(err)
	}
	fmt.Println(ans)
}

func TestNameFunc(t *testing.T) {
	str := "StudentTeacher"
	t.Log(nameFunc(str))
}
