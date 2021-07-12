package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Engine struct {
	db       *sql.DB
	nameFunc func(a string) string
}

func Open(dns string) (*Engine, error) {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	engine := &Engine{db, nameFunc}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return engine, nil
}

//修改命名方式
func (e *Engine) ChangeNameFunc(newFunc func(a string) string) {
	e.nameFunc = newFunc
}
