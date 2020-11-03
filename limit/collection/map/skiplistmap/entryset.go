package skiplistmap

import (
	"LimitGo/limit/collection"
	"reflect"
)

type EntrySet struct {
	sm *SkipListMap
}

// Size returns the number of elements in this collection.
func (ks *EntrySet) Size() int {
	return 0
}

// Empty returns true if this collection contains no element.
func (ks *EntrySet) Empty() bool {
	return false
}

// String returns a string representation of this collection.
func (ks *EntrySet) String() string {
	return ""
}

// Removes all of the elements from this collection.
func (ks *EntrySet) Clear() bool {
	return false
}

// GetIterator returns an iterator over the elements in this collection.
func (ks *EntrySet) GetIterator() collection.Itr {
	return nil
}

// Contains returns true if this list contains the specific element.
func (ks *EntrySet) Contains(p *collection.Object) bool {
	return false
}

// GetType returns type of the elements in this collection.
func (ks *EntrySet) GetType() reflect.Type {
	return nil
}

// Add inserts the specified element to this collection.
func (ks *EntrySet) Add(p *collection.Object) bool {
	return false
}

// Remove the first occurrence of the specified element from this collection.
func (ks *EntrySet) Remove(p *collection.Object) bool {
	return false
}

// AddAll appends all of the elements in the specified collection to this set.
func (ks *EntrySet) AddAll(list *collection.Linear) bool {
	return false
}

// Retains only the elements in this set that are contained in the
// specified collection.
func (ks *EntrySet) RetainAll(list *collection.Linear) bool {
	return false
}

// Removes from this set all of its elements that are contained in the
// specified collection.
func (ks *EntrySet) RemoveAll(list *collection.Linear) bool {
	return false
}

// Equals returns true only if the corresponding pairs of the elements
//in the two sets are equal.
func (ks *EntrySet) Equals(set *collection.Set) bool {
	return false
}
