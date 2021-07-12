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
