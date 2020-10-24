package linear

import (
	"LimitGo/limit/collection"
	"reflect"
)

// NewStack returns a new queue based on LinkedList.
func NewStack(t reflect.Type) *collection.Stack{
	var st collection.Stack = &LinkedList{nil, nil, 0, t}
	return &st
}
