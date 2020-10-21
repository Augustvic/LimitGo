package list

import (
	"LimitGo/limit/collections"
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

// ArrayList is one of the implementations of the List and is based on slice.
type ArrayList struct {
	elements []*collections.ListObject
	t reflect.Type
}

// ArrayListIterator represents the specific iterator of the ArrayList
type ArrayListIterator struct {
	List *ArrayList
	Cursor  int
	LastRet int
}

// NewArrayList returns a new arraylist.
func NewArrayList(t reflect.Type) *ArrayList {
	fmt.Println()
	l := ArrayList{}
	l.elements = make([]*collections.ListObject, 0, 10)
	l.t = t
	return &l
}

// Size returns the number of elements in this list.
func (l *ArrayList) Size() int {
	return len(l.elements)
}

// Empty returns true if this list contains no element.
func (l *ArrayList) Empty() bool {
	return len(l.elements) == 0
}

// Contains returns true if this list contains the specific element.
func (l *ArrayList) Contains(p *collections.ListObject) bool {
	if reflect.TypeOf(*p) != l.GetType() {
		return false
	}
	for _, v := range l.elements {
		if *v == *p {
			return true
		}
	}
	return false
}

// Append appends the specified element to the end of this list.
func (l *ArrayList) Append(p *collections.ListObject) {
	if reflect.TypeOf(*p) != l.GetType() {
		return
	}
	l.elements = append(l.elements, p)
}

// Insert the specified element at the specified position in this list.
func (l *ArrayList) Insert(index int, p *collections.ListObject) {
	if reflect.TypeOf(*p) != l.GetType() {
		return
	}
	copy(l.elements[index+1:], l.elements[index:])
	l.elements[index] = p
}

// Appends all of the elements in the specified list to the end of this list.
func (l *ArrayList) AddAll(list collections.List) {
	if list.GetType() != l.t {
		return
	}
	it := list.GetIterator()
	for it.HashNext() {
		v := it.Next()
		l.Append(v)
	}
}

// Remove the first occurrence of the specified element from this list.
//
func (l *ArrayList) Remove(p *collections.ListObject) {
	if reflect.TypeOf(*p) != l.GetType() {
		return
	}
	for i := 0; i < len(l.elements); i++ {
		if *p == *(l.elements[i]) {
			l.RemoveAt(i)
			return
		}
	}
}

// Removes the element at the specified position in this list.
func (l *ArrayList) RemoveAt(index int) {
	copy(l.elements[index:], l.elements[index+1:])
	l.elements = l.elements[:l.Size()-1]
}

// Removes all of the elements from this list.
func (l *ArrayList) Clear() {
	l.elements = l.elements[0:0]
}

// Equals returns true only if the corresponding pairs of the elements
// in the two lists are equal.
// Notice that equal means "==", not same address.
func (l *ArrayList) Equals(list collections.List) bool {
	if l.t != list.GetType() {
		return false
	}
	flag := l.Size() == list.Size()
	if !flag {
		return false
	}
	for i := 0; i < l.Size(); i++ {
		p1 := l.Get(i)
		p2 := list.Get(i)
		if *p1 != *p2 {
			return false
		}
	}
	return true
}

// String returns a string representation of this list.
func (l *ArrayList) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i := 0; i < len(l.elements); i++ {
		e := l.elements[i]
		if buf.Len() > len("{") {
			buf.WriteByte(',')
		}
		var s string
		if e == nil || *e == nil {
			s = "nil"
		} else {
			b, err := json.Marshal(*e)
			if err == nil {
				s = string(b)
			} else {
				s = "nil"
			}
		}
		_, err := fmt.Fprint(&buf, s)
		if err != nil {
			i--
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Get returns the element at the specified position in this list.
func (l *ArrayList) Get(index int) *collections.ListObject {
	if index < 0 || index >= l.Size() {
		panic("Index out of range!")
	}
	return l.elements[index]
}

// Set replaces the element at the specified position in this list with
//the specified element.
func (l *ArrayList) Set(index int, p *collections.ListObject) {
	if reflect.TypeOf(*p) != l.t || index >= l.Size() {
		return
	}
	l.elements[index] = p
}

// IndexOf returns the index of the first occurrence of the
//specified element
func (l *ArrayList) IndexOf(p *collections.ListObject) int {
	if reflect.TypeOf(*p) != l.t {
		return -1
	}
	for i := 0; i < l.Size(); i++ {
		if *p == *(l.elements[i]) {
			return i
		}
	}
	return -1
}

// GetIterator returns an iterator over the elements in this list.
func (l *ArrayList) GetIterator() collections.ListItr {
	 return &ArrayListIterator{l, 0, -1}
}

// GetType returns type of the elements in this list.
func (l *ArrayList) GetType() reflect.Type {
	return l.t
}

// HashNext returns true if the iteration has more elements.
func (it *ArrayListIterator) HashNext() bool {
	cursor := it.Cursor
	size := it.List.Size()
	return cursor != size
}

// Next returns the next element in the iteration.
func (it *ArrayListIterator) Next() *collections.ListObject {
	if it.HashNext() {
		lastRet := it.Cursor
		it.Cursor++
		return it.List.elements[lastRet]
	}
	return nil
}