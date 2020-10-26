package set

import (
	"LimitGo/limit/collection"
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

var Exists = struct{}{}

type Set struct {
	m map[collection.Object]struct{}
	t reflect.Type
}

// SetIterator represents the specific iterator of the Set.
// Note: SetIterator is just a snapshot of current elements.
type SetIterator struct {
	set *Set
	keys    []collection.Object
	cursor  int
	lastRet int
}

// New returns a new set.
func New(t reflect.Type) *Set {
	l := Set{make(map[collection.Object]struct{}), t}
	return &l
}

// Size returns the number of elements in this collection.
func (s *Set) Size() int {
	s.init()
	return len(s.m)
}

// Empty returns true if this collection contains no element.
func (s *Set) Empty() bool {
	s.init()
	return len(s.m) == 0
}

// GetIterator returns an iterator over the elements in this collection.
func (s *Set) GetIterator() collection.Itr {
	keys := getKeys(&s.m)
	it := SetIterator{s, keys, 0, -1}
	return &it
}

// GetType returns type of the elements in this collection.
func (s *Set) GetType() reflect.Type {
	return s.t
}

// String returns a string representation of this collection.
func (s *Set) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	it := s.GetIterator()
	for it.HashNext() {
		p := it.Next()
		if buf.Len() > len("{") {
			buf.WriteByte(',')
		}
		var s string
		b, err := json.Marshal(*p)
		if err == nil {
			s = string(b)
		} else {
			s = "nil"
		}
		_, err = fmt.Fprint(&buf, s)
	}
	buf.WriteByte('}')
	return buf.String()
}

// Removes all of the elements from this collection.
func (s *Set) Clear() bool {
	s.m = make(map[collection.Object]struct{})
	return true
}

// Contains returns true if this collection contains the specific element.
func (s *Set) Contains(p *collection.Object) bool {
	if s.checkNil(p) || !s.checkType(p) {
		return false
	}
	s.init()
	_, ok := s.m[*p]
	return ok
}

// Add inserts the specified element to this collection.
func (s *Set) Add(p *collection.Object) bool {
	if s.checkNil(p) || !s.checkType(p) {
		return false
	}
	s.init()
	s.m[*p] = Exists
	return true
}

// Remove the first occurrence of the specified element from this collection.
func (s *Set) Remove(p *collection.Object) bool {
	if s.checkNil(p) || !s.checkType(p) {
		return false
	}
	s.init()
	delete(s.m, *p)
	return true
}

// AddAll appends all of the elements in the specified collection to this set.
func (s *Set) AddAll(list *collection.Linear) bool {
	if list == nil || (*list) == nil || (*list).Empty() {
		return true
	}
	s.init()
	if s.t != (*list).GetType() {
		return false
	}
	it := (*list).GetIterator()
	for it.HashNext() {
		p := it.Next()
		s.Add(p)
	}
	return true
}

// Retains only the elements in this set that are contained in the
// specified collection.
func (s *Set) RetainAll(list *collection.Linear) bool {
	if list == nil || (*list) == nil || (*list).Empty() {
		return true
	}
	s.init()
	if s.t != (*list).GetType() {
		return false
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

// Removes from this set all of its elements that are contained in the
// specified collection.
func (s *Set) RemoveAll(list *collection.Linear) bool {
	if list == nil || (*list) == nil || (*list).Empty() {
		return true
	}
	s.init()
	if s.t != (*list).GetType() {
		return false
	}
	it := (*list).GetIterator()
	for it.HashNext() {
		s.Remove(it.Next())
	}
	return true
}

// Equals returns true only if the corresponding pairs of the elements
//in the two sets are equal.
func (s *Set) Equals(set *Set) bool {
	s.init()
	if set == nil || set.Size() != s.Size() {
		return false
	}
	it := set.GetIterator()
	for it.HashNext() {
		p := it.Next()
		if !s.Contains(p) {
			return false
		}
	}
	return true
}

// HashNext returns true if the iteration has more elements.
func (it *SetIterator) HashNext() bool {
	cursor := it.cursor
	size := len(it.keys)
	return cursor != size
}

// Next returns the next element in the iteration.
func (it *SetIterator) Next() *collection.Object {
	if it.HashNext() {
		it.lastRet = it.cursor
		it.cursor++
		return &it.keys[it.lastRet]
	}
	return nil
}

// Remove removes from the underlying collection the last element returned
// by this iterator.
func (it *SetIterator) Remove() (*collection.Object, bool) {
	if it.lastRet < 0 {
		return nil, false
	}
	key := it.keys[it.lastRet]
	copy(it.keys[it.lastRet:], it.keys[it.lastRet+1:])
	it.keys = it.keys[:len(it.keys)-1]
	it.set.Remove(&key)
	it.cursor = it.lastRet
	it.lastRet = -1
	return &key, true
}

func getKeys(m *map[collection.Object]struct{}) []collection.Object {
	keys := make([]collection.Object, 0, len(*m))
	for k := range *m {
		keys = append(keys, k)
	}
	return keys
}

func (s *Set) init() {
	if s.m == nil {
		s.m = make(map[collection.Object]struct{})
	}
}

func (s *Set) checkNil(p *collection.Object) bool {
	return p == nil || (*p) == nil
}

func (s *Set) checkIndex(index int) bool {
	return index >= 0 && index < s.Size()
}

func (s *Set) checkType(p *collection.Object) bool {
	return reflect.TypeOf(*p) == s.t
}