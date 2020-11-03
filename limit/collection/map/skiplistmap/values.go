package skiplistmap

import (
	"LimitGo/limit/collection"
	"reflect"
)

type Values struct {
	sm *SkipListMap
}

// Size returns the number of elements in this collection.
func (v *Values) Size() int {
	return 0
}

// Empty returns true if this collection contains no element.
func (v *Values) Empty() bool {
	return false
}

// String returns a string representation of this collection.
func (v *Values) String() string {
	return ""
}

// Removes all of the elements from this collection.
func (v *Values) Clear() bool {
	return false
}

// GetIterator returns an iterator over the elements in this collection.
func (v *Values) GetIterator() collection.Itr {
	return nil
}

// Contains returns true if this list contains the specific element.
func (v *Values) Contains(p *collection.Object) bool {
	return false
}

// GetType returns type of the elements in this collection.
func (v *Values) GetType() reflect.Type {
	return nil
}
