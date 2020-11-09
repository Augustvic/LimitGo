package skiplistmap

import (
	"LimitGo/limit/collection"
	"bytes"
	"encoding/json"
	"fmt"
)

type Values struct {
	sm *collection.SortedMap
}

type ValuesIterator struct {
	v *Values
	next    *Node
	lastRet *Node
}

// Size returns the number of elements in this collection.
func (v *Values) Size() int {
	return (*v.sm).Size()
}

// Empty returns true if this collection contains no element.
func (v *Values) Empty() bool {
	return (*v.sm).Size() == 0
}

// String returns a string representation of this collection.
func (v *Values) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	it := (*v.sm).GetEntryIterator()
	for it.HashNext() {
		entry := (*it.Next()).(collection.Entry)
		if buf.Len() > len("{") {
			buf.WriteByte(',')
		}
		value := entry.GetValue()
		var s string
		k, err1 := json.Marshal(*value)
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
func (v *Values) Clear() bool {
	return (*v.sm).Clear()
}

// GetIterator returns an iterator over the elements in this collection.
func (v *Values) GetIterator() collection.Itr {
	var m *SkipListMap
	var sm *SubMap
	var it collection.Itr
	m, ok := (*v.sm).(*SkipListMap)
	if !ok {
		sm = (*v.sm).(*SubMap)
	}
	if ok {
		it = &ValuesIterator{v, m.findFirst(), nil}
	} else {
		if sm != nil {
			it = &ValuesIterator{v, sm.loNode(), nil}
		}
	}
	return it
}

// Contains returns true if this list contains the specific element.
func (v *Values) Contains(p *collection.Object) bool {
	return (*v.sm).ContainsValue(p)
}

// HashNext returns true if the iteration has more elements.
func (it *ValuesIterator) HashNext() bool {
	return it.next != nil
}

// Next returns the next element in the iteration.
func (it *ValuesIterator) Next() *collection.Object {
	if it.HashNext() {
		it.lastRet = it.next
		it.next = it.next.next
		return it.lastRet.key
	}
	return nil
}

// Remove removes from the underlying collection the last element returned
// by this iterator.
func (it *ValuesIterator) Remove() (*collection.Object, bool) {
	return nil, false
}

// Unsupported operation.
func (v *Values) Add(p *collection.Object) bool {
	return false
}
