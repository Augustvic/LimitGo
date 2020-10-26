package arraylist

import (
	"LimitGo/limit/collection"
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

// ArrayList is one of the implementations of the List based on origin slice.
type ArrayList struct {
	elements []*collection.Object
	t        reflect.Type
}

// ArrayListIterator represents the specific iterator of the ArrayList
type ArrayListIterator struct {
	list    *ArrayList
	cursor  int
	lastRet int
}

// New returns a new ArrayList.
func New(t reflect.Type) *ArrayList {
	l := ArrayList{make([]*collection.Object, 0, 10), t}
	return &l
}

// NewArrayList returns a new ArrayList.
func NewArrayList(cap int, t reflect.Type) *ArrayList {
	l := ArrayList{make([]*collection.Object, 0, cap), t}
	return &l
}

// Size returns the number of elements in this list.
func (l *ArrayList) Size() int {
	l.checkInit()
	return len(l.elements)
}

// Empty returns true if this list contains no element.
func (l *ArrayList) Empty() bool {
	l.checkInit()
	return l.Size() == 0
}

// Contains returns true if this list contains the specific element.
func (l *ArrayList) Contains(p *collection.Object) bool {
	l.checkInit()
	if l.checkNil(p) || !l.checkType(p) {
		return false
	}
	for _, v := range l.elements {
		if reflect.DeepEqual(*v, *p) {
			return true
		}
	}
	return false
}

// Append appends the specified element to the end of this list.
func (l *ArrayList) Append(p *collection.Object) bool {
	l.checkInit()
	if l.checkNil(p) || !l.checkType(p) {
		return false
	}
	l.elements = append(l.elements, p)
	return true
}

// Insert the specified element at the specified position in this list.
func (l *ArrayList) Insert(index int, p *collection.Object) bool {
	l.checkInit()
	if l.checkNil(p) || !l.checkType(p) {
		return false
	}
	if index < 0 {
		index = 0
	}
	if index >= l.Size() {
		index = l.Size()
	}
	l.elements = append(l.elements, nil)
	if index < l.Size() {
		copy(l.elements[index+1:], l.elements[index:])
	}
	l.elements[index] = p
	return true
}

// AddAll appends all of the elements in the specified list to the end of this list.
func (l *ArrayList) AddAll(list *collection.Linear) bool {
	l.checkInit()
	if list == nil || *list == nil || (*list).Empty() {
		return true
	}
	if (*list).GetType() != l.t {
		return false
	}
	it := (*list).GetIterator()
	for it.HashNext() {
		v := it.Next()
		l.Append(v)
	}
	return true
}

// Remove the first occurrence of the specified element from this list.
func (l *ArrayList) Remove(p *collection.Object) bool {
	l.checkInit()
	if l.checkNil(p) {
		return true
	}
	if !l.checkType(p) {
		return false
	}
	for i := 0; i < len(l.elements); i++ {
		if reflect.DeepEqual(*p, *(l.elements[i])) {
			l.RemoveAt(i)
			return true
		}
	}
	return false
}

// Removes the element at the specified position in this list.
func (l *ArrayList) RemoveAt(index int) *collection.Object {
	l.checkInit()
	if !l.checkIndex(index) {
		return nil
	}
	p := l.elements[index]
	copy(l.elements[index:], l.elements[index+1:])
	l.elements = l.elements[:l.Size()-1]
	return p
}

// Removes all of the elements from this list.
func (l *ArrayList) Clear() bool {
	l.elements = l.elements[0:0]
	return true
}

// Equals returns true only if the corresponding pairs of the elements
// in the two lists are deep equal.
// Notice that equal do not means same address.
func (l *ArrayList) Equals(list *collection.List) bool {
	l.checkInit()
	if list == nil || *list == nil {
		return false
	}
	if l.t != (*list).GetType() {
		return false
	}
	if l.Size() != (*list).Size() {
		return false
	}
	for i := 0; i < l.Size(); i++ {
		p1 := l.Get(i)
		p2 := (*list).Get(i)
		if !reflect.DeepEqual(*p1, *p2) {
			return false
		}
	}
	return true
}

// String returns a string representation of this list.
func (l *ArrayList) String() string {
	l.checkInit()
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
func (l *ArrayList) Get(index int) *collection.Object {
	l.checkInit()
	if !l.checkIndex(index) {
		panic("Index out of range!")
	}
	return l.elements[index]
}

// Set replaces the element at the specified position in this list with
//the specified element.
func (l *ArrayList) Set(index int, p *collection.Object) bool {
	l.checkInit()
	if !l.checkIndex(index) || l.checkNil(p) || !l.checkType(p) || index >= l.Size() {
		return false
	}
	l.elements[index] = p
	return true
}

// IndexOf returns the index of the first occurrence of the
//specified element
func (l *ArrayList) IndexOf(p *collection.Object) int {
	l.checkInit()
	if l.checkNil(p) || reflect.TypeOf(*p) != l.t {
		return -1
	}
	for i := 0; i < l.Size(); i++ {
		if reflect.DeepEqual(*p, *(l.elements[i])) {
			return i
		}
	}
	return -1
}

// GetIterator returns an iterator over the elements in this list.
func (l *ArrayList) GetIterator() collection.Itr {
	l.checkInit()
	return &ArrayListIterator{l, 0, -1}
}

// GetType returns type of the elements in this list.
func (l *ArrayList) GetType() reflect.Type {
	l.checkInit()
	return l.t
}

// HashNext returns true if the iteration has more elements.
func (it *ArrayListIterator) HashNext() bool {
	cursor := it.cursor
	size := it.list.Size()
	return cursor != size
}

// Next returns the next element in the iteration.
func (it *ArrayListIterator) Next() *collection.Object {
	if it.HashNext() {
		it.lastRet = it.cursor
		it.cursor++
		return it.list.elements[it.lastRet]
	}
	return nil
}

// Remove removes from the underlying collection the last element returned
// by this iterator.
func (it *ArrayListIterator) Remove() (*collection.Object, bool) {
	if it.lastRet < 0 {
		return nil, false
	}
	p := it.list.RemoveAt(it.lastRet)
	it.cursor = it.lastRet
	it.lastRet = -1
	return p, true
}

// checkNil return true if p is nil or *p if nil
func (l *ArrayList) checkNil(p *collection.Object) bool {
	return p == nil || (*p) == nil
}

// checkIndex return true if index within the range
func (l *ArrayList) checkIndex(index int) bool {
	return index >= 0 && index < l.Size()
}

// checkType returns true if type matches
func (l *ArrayList) checkType(p *collection.Object) bool {
	return reflect.TypeOf(*p) == l.t
}

func (l *ArrayList) checkInit() {
	if l.elements == nil {
		l.elements = make([]*collection.Object, 0, 10)
	}
}
