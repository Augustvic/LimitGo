package deque

import (
	"LimitGo/limit/collection"
	"LimitGo/limit/collection/linear/linkedlist"
	"reflect"
)

// New returns a new deque based on LinkedList.
func New(t reflect.Type) *collection.Deque{
	var q collection.Deque = linkedlist.New(t)
	return &q
}
