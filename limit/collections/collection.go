package collections

import "LimitGo/limit"

type Iterator interface {
	HashNext() bool
	Next() *limit.Object
	Remove()
}