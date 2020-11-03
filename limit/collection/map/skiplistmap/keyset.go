package skiplistmap

import (
	"LimitGo/limit/collection"
	"reflect"
)

type KeySet struct {
	sm *SkipListMap
}

// Size returns the number of elements in this collection.
func (ks *KeySet) Size() int {
	return 0
}

// Empty returns true if this collection contains no element.
func (ks *KeySet) Empty() bool {
	return false
}

// String returns a string representation of this collection.
func (ks *KeySet) String() string {
	return ""
}

// Removes all of the elements from this collection.
func (ks *KeySet) Clear() bool {
	return false
}

// GetIterator returns an iterator over the elements in this collection.
func (ks *KeySet) GetIterator() collection.Itr {
	return nil
}

// Contains returns true if this list contains the specific element.
func (ks *KeySet) Contains(p *collection.Object) bool {
	return false
}

// GetType returns type of the elements in this collection.
func (ks *KeySet) GetType() reflect.Type {
	return nil
}

// Add inserts the specified element to this collection.
func (ks *KeySet) Add(p *collection.Object) bool {
	return false
}

// Remove the first occurrence of the specified element from this collection.
func (ks *KeySet) Remove(p *collection.Object) bool {
	return false
}

// AddAll appends all of the elements in the specified collection to this set.
func (ks *KeySet) AddAll(list *collection.Linear) bool {
	return false
}

// Retains only the elements in this set that are contained in the
// specified collection.
func (ks *KeySet) RetainAll(list *collection.Linear) bool {
	return false
}

// Removes from this set all of its elements that are contained in the
// specified collection.
func (ks *KeySet) RemoveAll(list *collection.Linear) bool {
	return false
}

// Equals returns true only if the corresponding pairs of the elements
//in the two sets are equal.
func (ks *KeySet) Equals(set *collection.Set) bool {
	return false
}

// SubSet returns a view of the portion of this set whose elements range from
// fromElement to toElement.  If fromElement and toElement are equal,
// the returned set is empty unless fromInclusive and toInclusive are both true.
func (ks *KeySet) SubSet(fromElement *collection.Object, fromInclusive bool, toElement *collection.Object, toInclusive bool) *collection.SortedSet {
	return nil
}

// HeadSet returns a view of the portion of this set whose elements are
// less than (or equal to, if inclusive is true) toElement
func (ks *KeySet) HeadSet(toElement *collection.Object, inclusive bool) *collection.SortedSet {
	return nil
}

// TailSet returns a view of the portion of this set whose elements are
// greater than (or equal to, if inclusive is true) fromElement.
func (ks *KeySet) TailSet(fromElement *collection.Object, inclusive bool) *collection.SortedSet {
	return nil
}

// First returns the first (lowest) element currently in this set.
func (ks *KeySet) First() *collection.Object {
	return nil
}

// Last returns the last (highest) element currently in this set.
func (ks *KeySet) Last() *collection.Object {
	return nil
}

// Lower returns the greatest element in this set strictly less than the given
// element, or null if there is no such element.
func (ks *KeySet) Lower(p *collection.Object) *collection.Object {
	return nil
}

// Floor returns the greatest element in this set less than or equal to
// the given element, or null if there is no such element.
func (ks *KeySet) Floor(p *collection.Object) *collection.Object {
	return nil
}

// Ceiling returns the least element in this set greater than or equal to
// the given element, or null if there is no such element.
func (ks *KeySet) Ceiling(p *collection.Object) *collection.Object {
	return nil
}

// Higher returns the least element in this set strictly greater than the
// given element, or null if there is no such element.
func (ks *KeySet) Higher(p *collection.Object) *collection.Object {
	return nil
}

// PollFirst retrieves and removes the first (lowest) element, or
// returns null if this set is empty.
func (ks *KeySet) PollFirst() *collection.Object {
	return nil
}

// PollLast retrieves and removes the last (highest) element, or
// returns null if this set is empty.
func (ks *KeySet) PollLast() *collection.Object {
	return nil
}

// DescendingSet returns a reverse order view of the elements contained
// in this set.
func (ks *KeySet) DescendingSet() *collection.SortedSet {
	return nil
}