package sortedset

import (
	"LimitGo/limit/collection"
	"bytes"
	"encoding/json"
	"fmt"
)

var Exists collection.Object = struct{}{}

// Sorted hashset based on SortedMap
type SortedSet struct {
	m *collection.SortedMap
}

func New(m *collection.SortedMap) *SortedSet {
	return &SortedSet{m}
}


// Size returns the number of elements in this collection.
func (s *SortedSet) Size() int {
	return (*s.m).Size()
}

// Empty returns true if this collection contains no element.
func (s *SortedSet) Empty() bool {
	return (*s.m).Size() == 0
}

// String returns a string representation of this collection.
func (s *SortedSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	it := (*s.m).GetEntryIterator()
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
func (s *SortedSet) Clear() bool {
	return (*s.m).Clear()
}

// GetIterator returns an iterator over the elements in this collection.
func (s *SortedSet) GetIterator() collection.Itr {
	return (*(*s.m).KeySet()).GetIterator()
}

// Contains returns true if this list contains the specific element.
func (s *SortedSet) Contains(p *collection.Object) bool {
	if s.checkNil(p) {
		return false
	}
	return (*s.m).ContainsKey(p)
}

// Add inserts the specified element to this collection.
func (s *SortedSet) Add(p *collection.Object) bool {
	if s.checkNil(p) {
		return false
	}
	ok, _ := (*s.m).Put(p, &Exists)
	return ok
}

// Remove the first occurrence of the specified element from this collection.
func (s *SortedSet) Remove(p *collection.Object) bool {
	if s.checkNil(p) {
		return false
	}
	(*s.m).Remove(p)
	return true
}

// AddAll appends all of the elements in the specified collection to this hashset.
func (s *SortedSet) AddAll(list *collection.Linear) bool {
	if list == nil || *list == nil || (*list).Size() == 0 {
		return true
	}
	it := (*list).GetIterator()
	for it.HashNext() {
		p := it.Next()
		(*s.m).Put(p, &Exists)
	}
	return true
}

// Retains only the elements in this hashset that are contained in the
// specified collection.
func (s *SortedSet) RetainAll(list *collection.Linear) bool {
	if list == nil || (*list) == nil || (*list).Empty() {
		s.Clear()
		return true
	}
	it := (*s).GetIterator()
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
func (s *SortedSet) RemoveAll(list *collection.Linear) bool {
	if list == nil || (*list) == nil || (*list).Empty() {
		return true
	}
	it := (*list).GetIterator()
	for it.HashNext() {
		s.Remove(it.Next())
	}
	return true
}

// Equals returns true only if the corresponding pairs of the elements
//in the two sets are equal.
func (s *SortedSet) Equals(set *collection.Set) bool {
	if set == nil || (*set).Size() != s.Size() {
		return false
	}
	it := (*set).GetIterator()
	for it.HashNext() {
		p := it.Next()
		if !s.Contains(p) {
			return false
		}
	}
	return true
}

// First returns the first (lowest) element currently in this hashset.
func (s *SortedSet) First() *collection.Object {
	entry := (*s.m).FirstEntry()
	return (*entry).GetKey()
}

// Last returns the last (highest) element currently in this hashset.
func (s *SortedSet) Last() *collection.Object {
	entry := (*s.m).LastEntry()
	return (*entry).GetKey()
}

// Lower returns the greatest element in this hashset strictly less than the given
// element, or null if there is no such element.
func (s *SortedSet) Lower(p *collection.Object) *collection.Object {
	entry := (*s.m).LowerEntry(p)
	if entry == nil {
		return nil
	} else {
		return (*entry).GetKey()
	}
}

// Floor returns the greatest element in this hashset less than or equal to
// the given element, or null if there is no such element.
func (s *SortedSet) Floor(p *collection.Object) *collection.Object {
	entry := (*s.m).FloorEntry(p)
	if entry == nil {
		return nil
	} else {
		return (*entry).GetKey()
	}
}

// Ceiling returns the least element in this hashset greater than or equal to
// the given element, or null if there is no such element.
func (s *SortedSet) Ceiling(p *collection.Object) *collection.Object {
	entry := (*s.m).CeilingEntry(p)
	if entry == nil {
		return nil
	} else {
		return (*entry).GetKey()
	}
}

// Higher returns the least element in this hashset strictly greater than the
// given element, or null if there is no such element.
func (s *SortedSet) Higher(p *collection.Object) *collection.Object {
	entry := (*s.m).HigherEntry(p)
	if entry == nil {
		return nil
	} else {
		return (*entry).GetKey()
	}
}

// PollFirst retrieves and removes the first (lowest) element, or
// returns null if this hashset is empty.
func (s *SortedSet) PollFirst() *collection.Object {
	entry := (*s.m).PollFirstEntry()
	if entry == nil {
		return nil
	} else {
		return (*entry).GetKey()
	}
}

// PollLast retrieves and removes the last (highest) element, or
// returns null if this hashset is empty.
func (s *SortedSet) PollLast() *collection.Object {
	entry := (*s.m).PollLastEntry()
	if entry == nil {
		return nil
	} else {
		return (*entry).GetKey()
	}
}

func (s *SortedSet) checkNil(p *collection.Object) bool {
	return p == nil || (*p) == nil
}
