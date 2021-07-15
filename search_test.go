package main

import (
	"fmt"
	"testing"

	"github.com/fffbbbbbb/reflact/table"
)

type TestTeacher struct {
	ID   int
	Name string
}

func TestSearchSlice(t *testing.T) {
	te := &[]table.TableDescInfo{}
	err := testDB.SearchSlice(te, "desc student")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", te)
}

func TestFind(t *testing.T) {
	ans := &[]StudentTest{}
	err := testDB.Column("ID", "Sex", "Name").Where("ID=2").Find(ans)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ans)
}

func TestFirst(t *testing.T) {
	ans := &StudentTest{}
	err := testDB.Column("ID", "Sex", "Name").Where("ID=8").First(ans)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ans)
}
