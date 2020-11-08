package skiplistmap

import (
	"LimitGo/limit/collection"
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

type EntrySet struct {
	sm *collection.SortedMap
}

type EntrySetIterator struct {
	es *EntrySet
	next    *Node
	lastRet *Node
}

// Size returns the number of elements in this collection.
func (es *EntrySet) Size() int {
	return (*es.sm).Size()
}

// Empty returns true if this collection contains no element.
func (es *EntrySet) Empty() bool {
	return (*es.sm).Size() == 0
}

// String returns a string representation of this collection.
func (es *EntrySet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	it := (*es.sm).GetEntryIterator()
	for it.HashNext() {
		entry := (*it.Next()).(collection.Entry)
		if buf.Len() > len("{") {
			buf.WriteByte(',')
		}
		key := entry.GetKey()
		value := entry.GetValue()
		var s string
		k, err1 := json.Marshal(*key)
		v, err2 := json.Marshal(*value)
		if err1 == nil && err2 == nil {
			s = string(k) + "=" + string(v)
		} else {
			s = "nil"
		}
		_, _ = fmt.Fprint(&buf, s)
	}
	buf.WriteByte('}')
	return buf.String()
}

// Removes all of the elements from this collection.
func (es *EntrySet) Clear() bool {
	return (*es.sm).Clear()
}

// GetIterator returns an iterator over the elements in this collection.
func (es *EntrySet) GetIterator() collection.Itr {
	sm := (*es.sm).(*SkipListMap)
	var it collection.Itr = &EntrySetIterator{es, sm.findFirst(), nil}
	return it
}

// Contains returns true if this list contains the specific element.
func (es *EntrySet) Contains(p *collection.Object) bool {
	entry := (*p).(collection.Entry)
	value := (*es.sm).Get(entry.GetKey())
	return value != nil && (*value) != nil && reflect.DeepEqual(entry.GetValue(), value)
}

// GetType returns type of the elements in this collection.
// Unsupported operation.
func (es *EntrySet) GetType() reflect.Type {
	return nil
}

// Add inserts the specified element to this collection.
// Unsupported operation.
func (es *EntrySet) Add(p *collection.Object) bool {
	return false
}

// Remove the first occurrence of the specified element from this collection.
// Unsupported operation.
func (es *EntrySet) Remove(p *collection.Object) bool {
	entry := (*p).(collection.Entry)
	(*es.sm).Remove(entry.GetKey())
	return true
}

// AddAll appends all of the elements in the specified collection to this hashset.
// Unsupported operation.
func (es *EntrySet) AddAll(list *collection.Linear) bool {
	return false
}

// Retains only the elements in this hashset that are contained in the
// specified collection.
// Unsupported operation.
func (es *EntrySet) RetainAll(list *collection.Linear) bool {
	if list == nil || (*list) == nil || (*list).Empty() {
		return true
	}
	it := (*es).GetIterator()
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
// Unsupported operation.
func (es *EntrySet) RemoveAll(list *collection.Linear) bool {
	if list == nil || (*list) == nil || (*list).Empty() {
		return true
	}
	it := (*list).GetIterator()
	for it.HashNext() {
		es.Remove(it.Next())
	}
	return true
}

// Equals returns true only if the corresponding pairs of the elements
//in the two sets are equal.
func (es *EntrySet) Equals(set *collection.Set) bool {
	if set == nil || (*set).Size() != es.Size() {
		return false
	}
	it := (*set).GetIterator()
	for it.HashNext() {
		p := it.Next()
		if !es.Contains(p) {
			return false
		}
	}
	return true
}


// HashNext returns true if the iteration has more elements.
func (it *EntrySetIterator) HashNext() bool {
	return it.next != nil
}

// Next returns the next element in the iteration.
func (it *EntrySetIterator) Next() *collection.Object {
	if it.HashNext() {
		it.lastRet = it.next
		it.next = it.next.next
		return it.lastRet.key
	}
	return nil
}

// Remove removes from the underlying collection the last element returned
// by this iterator.
func (it *EntrySetIterator) Remove() (*collection.Object, bool) {
	if it.lastRet == nil {
		return nil, false
	}
	var last collection.Entry = it.lastRet
	lastNext := it.lastRet.next
	(*it.es.sm).Remove(last.GetKey())
	if it.next == it.lastRet {
		it.next = lastNext
	}
	it.lastRet = nil
	var obj collection.Object = last
	return &obj, true
}
