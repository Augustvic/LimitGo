package stack

import (
	"LimitGo/limit/collection"
	"LimitGo/limit/collection/linear/linkedlist"
	"reflect"
)

// New returns a new queue based on LinkedList.
func New(t reflect.Type) *collection.Stack{
	var st collection.Stack = linkedlist.NewLinkedList(nil, nil, 0, t)
	return &st
}
