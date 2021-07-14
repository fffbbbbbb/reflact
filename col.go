package main

func (e *Engine) getInstance() *Engine {
	return &Engine{
		db:        e.db,
		nameFunc:  e.nameFunc,
		DBVersion: e.DBVersion,
		hasJson:   e.hasJson,
		column:    e.column,
		where:     e.where,
	}
}

func (e *Engine) Column(args ...string) *Engine {
	tx := e.getInstance()
	tx.column = args
	return tx
}

func (e *Engine) Where(args string) *Engine {
	tx := e.getInstance()
	tx.where = args
	return tx
}
