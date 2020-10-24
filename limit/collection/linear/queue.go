package linear

import (
	"LimitGo/limit/collection"
	"reflect"
)

// NewQueue returns a new queue based on LinkedList.
func NewQueue(t reflect.Type) *collection.Queue{
	var q collection.Queue = &LinkedList{nil, nil, 0, t}
	return &q
}

// NewDeque returns a new deque based on LinkedList.
func NewDeque(t reflect.Type) *collection.Deque{
	var q collection.Deque = &LinkedList{nil, nil, 0, t}
	return &q
}
