package main

import (
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	username := "root"
	password := "123456"
	address := "127.0.0.1"
	dbname := "relact"
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username, password, address, dbname)
	if _, err := Open(dns); err != nil {
		t.Log(err)
	}
}
