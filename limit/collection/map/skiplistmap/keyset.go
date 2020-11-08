package skiplistmap

import (
	"LimitGo/limit/collection"
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

type KeySet struct {
	sm *collection.SortedMap
}

type KeySetIterator struct {
	ks *KeySet
	next    *Node
	lastRet *Node
}

// Size returns the number of elements in this collection.
func (ks *KeySet) Size() int {
	return (*ks.sm).Size()
}

// Empty returns true if this collection contains no element.
func (ks *KeySet) Empty() bool {
	return (*ks.sm).Size() == 0
}

// String returns a string representation of this collection.
func (ks *KeySet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	it := (*ks.sm).GetEntryIterator()
	for it.HashNext() {
		entry := (*it.Next()).(collection.Entry)
		if buf.Len() > len("{") {
			buf.WriteByte(',')
		}
		key := entry.GetKey()
		var s string
		k, err1 := json.Marshal(*key)
		if err1 == nil {
			s = string(k)
		} else {
			s = "nil"
		}
		_, _ = fmt.Fprint(&buf, s)
	}
	buf.WriteByte('}')
	return buf.String()
}

// Removes all of the elements from this collection.
func (ks *KeySet) Clear() bool {
	return (*ks.sm).Clear()
}

// GetIterator returns an iterator over the elements in this collection.
func (ks *KeySet) GetIterator() collection.Itr {
	sm := (*ks.sm).(*SkipListMap)
	var it collection.Itr = &KeySetIterator{ks, sm.findFirst(), nil}
	return it
}

// Contains returns true if this list contains the specific element.
func (ks *KeySet) Contains(p *collection.Object) bool {
	return (*ks.sm).ContainsKey(p)
}

// GetType returns type of the elements in this collection.
func (ks *KeySet) GetType() reflect.Type {
	return (*ks.sm).GetKeyType()
}

// Add inserts the specified element to this collection.
// Unsupported operation.
func (ks *KeySet) Add(p *collection.Object) bool {
	return false
}

// Remove the first occurrence of the specified element from this collection.
func (ks *KeySet) Remove(p *collection.Object) bool {
	(*ks.sm).Remove(p)
	return true
}

// AddAll appends all of the elements in the specified collection to this hashset.
// Unsupported operation.
func (ks *KeySet) AddAll(list *collection.Linear) bool {
	return false
}

// Retains only the elements in this hashset that are contained in the
// specified collection.
func (ks *KeySet) RetainAll(list *collection.Linear) bool {
	if list == nil || (*list) == nil || (*list).Empty() {
		return true
	}
	if (*ks.sm).GetKeyType() != (*list).GetType() {
		return false
	}
	it := (*ks).GetIterator()
	for it.HashNext() {
		p := it.Next()
		if !(*list).Contains(p) {
			it.Remove()
		}
	}
	return true
}

// Removes from this hashset all of its elements that are contained in the
// specified collection.
func (ks *KeySet) RemoveAll(list *collection.Linear) bool {
	if list == nil || (*list) == nil || (*list).Empty() {
		return true
	}
	if (*ks.sm).GetKeyType() != (*list).GetType() {
		return false
	}
	it := (*list).GetIterator()
	for it.HashNext() {
		ks.Remove(it.Next())
	}
	return true
}

// Equals returns true only if the corresponding pairs of the elements
//in the two sets are equal.
func (ks *KeySet) Equals(set *collection.Set) bool {
	if set == nil || (*set).Size() != ks.Size() {
		return false
	}
	it := (*set).GetIterator()
	for it.HashNext() {
		p := it.Next()
		if !ks.Contains(p) {
			return false
		}
	}
	return true
}

// SubSet returns a view of the portion of this hashset whose elements range from
// fromElement to toElement.  If fromElement and toElement are equal,
// the returned hashset is empty unless fromInclusive and toInclusive are both true.
func (ks *KeySet) SubSet(fromElement *collection.Object, fromInclusive bool, toElement *collection.Object, toInclusive bool) *collection.SortedSet {
	if fromElement != nil && !ks.checkType(fromElement) {
		return nil
	}
	if toElement != nil && !ks.checkType(toElement) {
		return nil
	}
	sm := (*ks.sm).(*SkipListMap)
	p := SubMap{sm, fromElement, toElement, fromInclusive, toInclusive, nil, nil, nil}
	var t collection.SortedMap = &p
	var ret collection.SortedSet = &KeySet{&t}
	return &ret
}

// HeadSet returns a view of the portion of this hashset whose elements are
// less than (or equal to, if inclusive is true) toElement
func (ks *KeySet) HeadSet(toElement *collection.Object, inclusive bool) *collection.SortedSet {
	if !ks.checkType(toElement) {
		return nil
	}
	return ks.SubSet(nil, false, toElement, inclusive)
}

// TailSet returns a view of the portion of this hashset whose elements are
// greater than (or equal to, if inclusive is true) fromElement.
func (ks *KeySet) TailSet(fromElement *collection.Object, inclusive bool) *collection.SortedSet {
	if !ks.checkType(fromElement) {
		return nil
	}
	return ks.SubSet(fromElement, inclusive, nil, false)
}

// First returns the first (lowest) element currently in this hashset.
func (ks *KeySet) First() *collection.Object {
	 entry := (*ks.sm).FirstEntry()
	 return (*entry).GetKey()
}

// Last returns the last (highest) element currently in this hashset.
func (ks *KeySet) Last() *collection.Object {
	entry := (*ks.sm).LastEntry()
	return (*entry).GetKey()
}

// Lower returns the greatest element in this hashset strictly less than the given
// element, or null if there is no such element.
func (ks *KeySet) Lower(p *collection.Object) *collection.Object {
	entry := (*ks.sm).LowerEntry(p)
	if entry == nil {
		return nil
	} else {
		return (*entry).GetKey()
	}
}

// Floor returns the greatest element in this hashset less than or equal to
// the given element, or null if there is no such element.
func (ks *KeySet) Floor(p *collection.Object) *collection.Object {
	entry := (*ks.sm).FloorEntry(p)
	if entry == nil {
		return nil
	} else {
		return (*entry).GetKey()
	}
}

// Ceiling returns the least element in this hashset greater than or equal to
// the given element, or null if there is no such element.
func (ks *KeySet) Ceiling(p *collection.Object) *collection.Object {
	entry := (*ks.sm).CeilingEntry(p)
	if entry == nil {
		return nil
	} else {
		return (*entry).GetKey()
	}
}

// Higher returns the least element in this hashset strictly greater than the
// given element, or null if there is no such element.
func (ks *KeySet) Higher(p *collection.Object) *collection.Object {
	entry := (*ks.sm).HigherEntry(p)
	if entry == nil {
		return nil
	} else {
		return (*entry).GetKey()
	}
}

// PollFirst retrieves and removes the first (lowest) element, or
// returns null if this hashset is empty.
func (ks *KeySet) PollFirst() *collection.Object {
	entry := (*ks.sm).PollFirstEntry()
	if entry == nil {
		return nil
	} else {
		return (*entry).GetKey()
	}
}

// PollLast retrieves and removes the last (highest) element, or
// returns null if this hashset is empty.
func (ks *KeySet) PollLast() *collection.Object {
	entry := (*ks.sm).PollLastEntry()
	if entry == nil {
		return nil
	} else {
		return (*entry).GetKey()
	}
}

// HashNext returns true if the iteration has more elements.
func (it *KeySetIterator) HashNext() bool {
	return it.next != nil
}

// Next returns the next element in the iteration.
func (it *KeySetIterator) Next() *collection.Object {
	if it.HashNext() {
		it.lastRet = it.next
		it.next = it.next.next
		return it.lastRet.key
	}
	return nil
}

// Remove removes from the underlying collection the last element returned
// by this iterator.
func (it *KeySetIterator) Remove() (*collection.Object, bool) {
	if it.lastRet == nil {
		return nil, false
	}
	var last collection.Entry = it.lastRet
	lastNext := it.lastRet.next
	(*it.ks.sm).Remove(last.GetKey())
	if it.next == it.lastRet {
		it.next = lastNext
	}
	it.lastRet = nil
	return last.GetKey(), true
}

func (ks *KeySet) checkType(p *collection.Object) bool {
	return reflect.TypeOf(*p) == (*ks.sm).GetKeyType()
}