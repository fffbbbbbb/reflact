package errinfo

import "errors"

var (
	KindNoSruct = errors.New("param is not struct")
	KindNoSlice = errors.New("param is not slice")
)
