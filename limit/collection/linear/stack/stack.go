package stack

import (
	"LimitGo/limit/collection"
	"LimitGo/limit/collection/linear/linkedlist"
	"reflect"
)

// New returns a new stack based on LinkedList.
func New(t reflect.Type) *collection.Stack{
	var st collection.Stack = linkedlist.New(t)
	return &st
}
