package priorityqueue

import (
	"LimitGo/limit/collection"
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

const initCap = 8

type PriorityQueue struct {
	elements []*collection.Object
	precede func(p1 *collection.Object, p2 *collection.Object) bool
}

// PriorityQueueIterator represents the specific iterator of the PriorityQueue.
type Iterator struct {
	list    *PriorityQueue
	cursor  int
	lastRet int
}

func New(precede func(p1 *collection.Object, p2 *collection.Object) bool) *PriorityQueue {
	return &PriorityQueue{make([]*collection.Object, 0, initCap), precede}
}

func (q *PriorityQueue) GetFunc() func(p1 *collection.Object, p2 *collection.Object) bool {
	return q.precede
}

// Size returns the number of elements in this collection.
func (q *PriorityQueue) Size() int {
	return len(q.elements)
}

// Empty returns true if this collection contains no element.
func (q *PriorityQueue) Empty() bool {
	return q.Size() == 0
}

// GetIterator returns an iterator over the elements in this collection.
func (q *PriorityQueue) GetIterator() collection.Itr {
	return &Iterator{q, 0, -1}
}

// String returns a string representation of this collection.
func (q *PriorityQueue) String() string {
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
func (q *PriorityQueue) Clear() bool {
	q.elements = make([]*collection.Object, 0, initCap)
	return true
}

// Contains returns true if this collection contains the specific element.
func (q *PriorityQueue) Contains(p *collection.Object) bool {
	if q.checkNil(p) {
		return false
	}
	for _, v := range q.elements {
		if reflect.DeepEqual(*v, *p) {
			return true
		}
	}
	return false
}

// Peek returns the head of this queue, or nil if this queue is empty
func (q *PriorityQueue) First() *collection.Object {
	if q.Size() == 0 {
		return nil
	} else {
		return q.elements[0]
	}
}

// Poll returns and removes the head of this queue, or nil if this queue is empty
func (q *PriorityQueue) Poll() *collection.Object {
	if q.Size() == 0 {
		return nil
	}
	result := q.elements[0]
	x := q.elements[len(q.elements) - 1]
	q.elements = q.elements[0:len(q.elements)-1]
	if len(q.elements) != 0 {
		q.siftDown(0, x)
	}
	return result
}

// Add inserts the specified element to the end of this queue.
func (q *PriorityQueue) Add(p *collection.Object) bool {
	if q.checkNil(p) {
		return false
	}
	if q.Size() == 0 {
		q.elements = append(q.elements, p)
		return true
	}
	q.elements = append(q.elements, nil)
	q.siftUp(len(q.elements) - 1, p)
	return true
}

// checkNil return true if p is nil or *p if nil
func (q *PriorityQueue) checkNil(p *collection.Object) bool {
	return p == nil || (*p) == nil
}

// checkIndex return true if index within the range
func (q *PriorityQueue) checkIndex(index int) bool {
	return index >= 0 && index < q.Size()
}

// HashNext returns true if the iteration has more elements.
func (it *Iterator) HashNext() bool {
	cursor := it.cursor
	size := it.list.Size()
	return cursor != size
}

// Next returns the next element in the iteration.
func (it *Iterator) Next() *collection.Object {
	if it.HashNext() {
		it.lastRet = it.cursor
		it.cursor++
		return it.list.elements[it.lastRet]
	}
	return nil
}

// Remove removes from the underlying collection the last element returned
// by this iterator.
func (it *Iterator) Remove() (*collection.Object, bool) {
	if it.lastRet < 0 {
		return nil, false
	}
	p := it.list.removeAt(it.lastRet)
	it.cursor = it.lastRet
	it.lastRet = -1
	return p, true
}

func (q *PriorityQueue) removeAt(index int) *collection.Object {
	if !q.checkIndex(index) {
		return nil
	}
	p := q.elements[index]
	copy(q.elements[index:], q.elements[index+1:])
	q.elements = q.elements[:q.Size()-1]
	return p
}

func (q *PriorityQueue) siftDown(k int, x *collection.Object) {
	half := q.Size() >> 1
	for k < half {
		child := k << 1 + 1
		p := q.elements[child]
		right := child + 1
		if right < q.Size() && q.precede(q.elements[right], p) {
			child = right
			p = q.elements[child]
		}
		if q.precede(x, p) {
			break
		}
		q.elements[k] = p
		k = child
	}
	q.elements[k] = x
}

func (q *PriorityQueue) siftUp(k int, x *collection.Object) {
	for k > 0 {
		parent := (k - 1) >> 1
		e := q.elements[parent]
		if q.precede(e, x) {
			break
		}
		q.elements[k] = e
		k = parent
	}
	q.elements[k] = x
}