package main

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/fffbbbbbb/reflact/errinfo"

	"github.com/fffbbbbbb/reflact/table"

	_ "github.com/go-sql-driver/mysql"
)

type Engine struct {
	db        *sql.DB
	nameFunc  func(a string) string
	DBVersion string
	hasJson   bool
	column    []string
	where     string
}

func Open(dns string) (*Engine, error) {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	engine := &Engine{
		db:       db,
		nameFunc: nameFunc,
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	ans := []table.Version{}
	err = engine.SearchSlice(&ans, "select version() as Version")
	if err != nil {
		return nil, errinfo.GetVersionError
	}
	engine.DBVersion = ans[0].Version
	engine.changeHasJson(hasJsonByVersion(engine.DBVersion))
	return engine, nil
}

//修改命名方式
func (e *Engine) ChangeNameFunc(newFunc func(a string) string) {
	e.nameFunc = newFunc
}

func (e *Engine) changeHasJson(a bool) {
	e.hasJson = a
}

func hasJsonByVersion(version string) bool {
	hasJson := false
	if version != "" {
		versionSlice := strings.Split(version, ".")
		v0, _ := strconv.Atoi(versionSlice[0])
		if v0 > 5 {
			hasJson = true
		} else if v0 == 5 && len(versionSlice) > 1 {
			v1, _ := strconv.Atoi(versionSlice[1])
			if v1 >= 7 {
				hasJson = true
			}
		}
	}
	return hasJson
}
