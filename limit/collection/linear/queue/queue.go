package queue

import (
	"LimitGo/limit/collection"
	"LimitGo/limit/collection/linear/linkedlist"
	"reflect"
)

// New returns a new queue based on LinkedList.
func New(t reflect.Type) *collection.Queue{
	var q collection.Queue = linkedlist.New(t)
	return &q
}