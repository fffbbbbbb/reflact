package main

import (
	"github.com/fffbbbbbb/reflact/table"
)

func (db *Engine) SyncTable(opt ...interface{}) error {
	for _, v := range opt {
		err := table.SyncTable(db.db, v)
		if err != nil {
			return err
		}
	}
	return nil
}
