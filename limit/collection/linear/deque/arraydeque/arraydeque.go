package arraydeque

import (
	"LimitGo/limit/collection"
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

const initLen int = 8

type ArrayDeque struct {
	elements []*collection.Object
	head     int
	tail     int
	t        reflect.Type
}

// ArrayDequeIterator represents the specific iterator of the ArrayDeque
type ArrayDequeIterator struct {
	list    *ArrayDeque
	cursor  int
	lastRet int
	fence   int
}

func New(t reflect.Type) *ArrayDeque {
	l := ArrayDeque{make([]*collection.Object, initLen, initLen), 0, 0, t}
	return &l
}

// Size returns the number of elements in this collection.
func (q *ArrayDeque) Size() int {
	return (q.tail - q.head) & (len(q.elements) - 1)
}

// Empty returns true if this collection contains no element.
func (q *ArrayDeque) Empty() bool {
	return q.head == q.tail
}

// GetIterator returns an iterator over the elements in this collection.
func (q *ArrayDeque) GetIterator() collection.Itr {
	return &ArrayDequeIterator{q, q.head, -1, q.tail}
}

// GetType returns type of the elements in this collection.
func (q *ArrayDeque) GetType() reflect.Type {
	return q.t
}

// String returns a string representation of this collection.
func (q *ArrayDeque) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	it := q.GetIterator()
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
func (q *ArrayDeque) Clear() bool {
	for i := 0; i < len(q.elements); i++ {
		q.elements[i] = nil
	}
	q.head = 0
	q.tail = 0
	return true
}

// Contains returns true if this collection contains the specific element.
func (q *ArrayDeque) Contains(p *collection.Object) bool {
	q.checkInit()
	if q.checkNil(p) || !q.checkType(p) {
		return false
	}
	mask := len(q.elements) - 1
	i := q.head
	v := q.elements[i]
	for i != q.tail {
		if reflect.DeepEqual(*v, *p) {
			return true
		}
		i = (i + 1) & mask
		v = q.elements[i]
	}
	return false
}

// AddFirst inserts the specified element at the front of this deque.
func (q *ArrayDeque) AddFirst(p *collection.Object) bool {
	q.checkInit()
	if q.checkNil(p) || !q.checkType(p) {
		return false
	}
	q.head = (q.head - 1) & (len(q.elements) - 1)
	q.elements[q.head] = p
	if q.head == q.tail {
		q.doubleLen()
	}
	return true
}

// AddLast inserts the specified element at the end of this deque.
func (q *ArrayDeque) AddLast(p *collection.Object) bool {
	q.checkInit()
	if q.checkNil(p) || !q.checkType(p) {
		return false
	}
	q.elements[q.tail] = p
	q.tail = (q.tail + 1) & (len(q.elements) - 1)
	if q.head == q.tail {
		q.doubleLen()
	}
	return true
}

// RemoveFirst removes and returns the head of this deque,
// or returns nil if this deque is empty.
func (q *ArrayDeque) RemoveFirst() *collection.Object {
	q.checkInit()
	h := q.head
	result := q.elements[h]
	if result == nil {
		return nil
	}
	q.elements[h] = nil
	q.head = (h + 1) & (len(q.elements) - 1)
	return result
}

// RemoveLast removes and returns the tail of this deque,
// or returns nil if this deque is empty.
func (q *ArrayDeque) RemoveLast() *collection.Object {
	q.checkInit()
	t := (q.tail - 1) & (len(q.elements) - 1)
	result := q.elements[t]
	if result == nil {
		return nil
	}
	q.elements[t] = nil
	q.tail = t
	return result
}

// GetFirst returns the head of this queue, or nil if this deque is empty.
func (q *ArrayDeque) GetFirst() *collection.Object {
	q.checkInit()
	return q.elements[q.head]
}

// GetLast returns the tail of this queue, or nil if this deque is empty.
func (q *ArrayDeque) GetLast() *collection.Object {
	q.checkInit()
	return q.elements[(q.tail-1)&(len(q.elements)-1)]
}

// HashNext returns true if the iteration has more elements.
func (it *ArrayDequeIterator) HashNext() bool {
	return it.cursor != it.fence
}

// Next returns the next element in the iteration.
func (it *ArrayDequeIterator) Next() *collection.Object {
	if it.HashNext() {
		result := it.list.elements[it.cursor]
		it.lastRet = it.cursor
		it.cursor = (it.cursor + 1) & (len(it.list.elements) - 1)
		return result
	}
	return nil
}

// Remove removes from the underlying collection the last element returned
// by this iterator.
func (it *ArrayDequeIterator) Remove() (*collection.Object, bool) {
	if it.lastRet < 0 {
		return nil, false
	}
	p := it.list.removeAt(it.lastRet)
	if p != nil {
		it.cursor = (it.cursor - 1) & (len(it.list.elements) - 1)
		it.fence = it.list.tail
		it.lastRet = -1
		return p, true
	} else {
		return p, false
	}
}

func (q *ArrayDeque) removeAt(index int) *collection.Object {
	q.checkInit()
	removed := q.elements[index]
	if !q.checkIndex(index) || removed == nil {
		return nil
	}
	elements := q.elements
	mask := len(elements) - 1
	h := q.head
	t := q.tail
	front := (index - h) & mask
	back := (t - index) & mask
	if front < back {
		if h < index {
			copy(elements[h+1:index+1], elements[h:index])
		} else {
			copy(elements[1:index+1], elements[0:index])
			elements[0] = elements[mask]
			copy(elements[h+1:mask+1],elements[h:mask])
		}
		elements[h] = nil
		q.head = (h + 1) & mask
	} else {
		if index < t {
			copy(elements[index:t-1], elements[index+1:t])
			q.tail = t - 1
		} else {
			copy(elements[index:mask], elements[index+1:mask+1])
			elements[mask] = elements[0]
			copy(elements[0:t-1], elements[1:t])
			q.tail = (t - 1) & mask
		}
		elements[q.tail] = nil
	}
	return removed
}

// checkNil return true if p is nil or *p if nil
func (q *ArrayDeque) checkNil(p *collection.Object) bool {
	return p == nil || (*p) == nil
}

// checkIndex return true if index within the range
func (q *ArrayDeque) checkIndex(index int) bool {
	if q.head == q.tail {
		return false
	}
	if index >= len(q.elements) || index < 0 {
		return false
	}
	if q.head < q.tail {
		return (index >= q.head) && (index < q.tail)
	} else {
		return index > q.head || index < q.tail
	}
}

// checkType returns true if type matches
func (q *ArrayDeque) checkType(p *collection.Object) bool {
	return reflect.TypeOf(*p) == q.t
}

func (q *ArrayDeque) checkInit() {
	if q.elements == nil {
		q.elements = make([]*collection.Object, initLen, initLen)
	}
}

func (q *ArrayDeque) doubleLen() {
	newLen := 2 * len(q.elements)
	if newLen < 0 {
		panic("Sorry, deque too big")
	}
	slice := make([]*collection.Object, newLen, newLen)
	p := q.head
	n := len(q.elements)
	r := n - p
	copy(slice[0:r], q.elements[p:n])
	copy(slice[r:n], q.elements[0:p])
	q.elements = slice
	q.head = 0
	q.tail = n
}
