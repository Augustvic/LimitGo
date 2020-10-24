package linear

import (
	"LimitGo/limit/collection"
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

type Node struct {
	item *collection.LinearObject
	prev *Node
	next *Node
}

// Origin list in Go SDK see package o"container/list"
// LinkedList (exactly, every collection) do not allow existence of two different types.
type LinkedList struct {
	head *Node
	tail *Node
	size int
	t reflect.Type
}

// LinkedListIterator represents the specific iterator of the LinkedList
type LinkedListIterator struct {
	list *LinkedList
	next *Node
	lastRet *Node
}

// NewArrayList returns a new LinkedList.
func NewLinkedList(t reflect.Type) *LinkedList{
	l := LinkedList{nil, nil, 0, t}
	return &l
}


// Size returns the number of elements in this collection.
func (l *LinkedList) Size() int {
	return l.size
}

// Empty returns true if this collection contains no element.
func (l *LinkedList) Empty() bool {
	return l.size == 0
}

// GetIterator returns an iterator over the elements in this collection.
func (l *LinkedList) GetIterator() collection.LinearItr {
	return &LinkedListIterator{l, l.head, nil}
}

// GetType returns type of the elements in this collection.
func (l *LinkedList) GetType() reflect.Type {
	return l.t
}

// String returns a string representation of this collection.
func (l *LinkedList) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	it := l.GetIterator()
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
func (l *LinkedList) Clear() bool {
	l.size = 0
	l.head = nil
	l.tail = nil
	return true
}

// Contains returns true if this list contains the specific element.
func (l *LinkedList) Contains(p *collection.LinearObject) bool {
	if l.checkNil(p) || !l.checkType(p) {
		return false
	}
	return l.IndexOf(p) != -1
}

// Append appends the specified element to the end of this list.
func (l *LinkedList) Append(p *collection.LinearObject) bool {
	if l.checkNil(p) || !l.checkType(p) {
		return false
	}
	return l.linkLast(p)
}

// Insert the specified element at the specified position in this list.
func (l *LinkedList) Insert(index int, p *collection.LinearObject) bool {
	if l.checkNil(p) || !l.checkType(p) {
		return false
	}
	if index <= 0 {
		l.linkFirst(p)
	} else if index >= l.size {
		l.linkLast(p)
	} else {
		l.linkBefore(p, l.node(index))
	}
	return true
}

// AddAll appends all of the elements in the specified list to the end of this list.
func (l *LinkedList) AddAll(list *collection.Linear) bool {
	if list == nil || (*list) == nil || (*list).Empty() {
		return true
	}
	if l.t != (*list).GetType() {
		return false
	}
	it := (*list).GetIterator()
	for it.HashNext() {
		p := it.Next()
		l.linkLast(p)
	}
	return true
}

// AddAllHead appends all of the elements in the specified list to the start of this list.
func (l *LinkedList) AddAllHead(list *collection.Linear) bool {
	if list == nil || (*list) == nil || (*list).Empty() {
		return true
	}
	if l.t != (*list).GetType() {
		return false
	}
	var p *Node = nil
	var q *Node = nil
	it := (*list).GetIterator()
	for it.HashNext() {
		e := it.Next()
		if p == nil || q == nil {
			node := &Node{e, nil, nil}
			p = node
			q = node
		} else {
			node := &Node{e, q, nil}
			q.next = node
			q = q.next
		}
	}
	if p != nil && q != nil {
		if l.Empty() {
			l.head = p
			l.tail = q
		} else {
			q.next = l.head
			l.head.prev = q
			l.head = p

		}
	}
	l.size += (*list).Size()
	return true
}

// Remove the first occurrence of the specified element from this list.
func (l *LinkedList) Remove(p *collection.LinearObject) bool {
	if l.checkNil(p) {
		return true
	}
	if !l.checkType(p) {
		return false
	}
	return l.unlink(l.nodeBy(p))
}

// Removes the element at the specified position in this list.
func (l *LinkedList) RemoveAt(index int) *collection.LinearObject {
	if !l.checkIndex(index) {
		return nil
	}
	p := l.node(index)
	if l.unlink(p) {
		return p.item
	} else {
		return nil
	}
}

// Equals returns true only if the corresponding pairs of the elements
//in the two lists are equal.
func (l *LinkedList) Equals(list *collection.List) bool {
	if list == nil || *list == nil {
		return false
	}
	if l.t != (*list).GetType() {
		return false
	}
	if l.Size() != (*list).Size() {
		return false
	}
	it1 := l.GetIterator()
	it2 := (*list).GetIterator()
	for it1.HashNext() && it2.HashNext() {
		p1 := it1.Next()
		p2 := it2.Next()
		if !reflect.DeepEqual(*p1, *p2) {
			return false
		}
	}
	return true
}

// Get returns the element at the specified position in this list.
func (l *LinkedList) Get(index int) *collection.LinearObject {
	if !l.checkIndex(index) {
		return nil
	}
	return l.node(index).item
}

// Set replaces the element at the specified position in this list with
//the specified element.
func (l *LinkedList) Set(index int, p *collection.LinearObject) bool {
	if !l.checkIndex(index) || l.checkNil(p) || !l.checkType(p) {
		return false
	}
	node := l.node(index)
	node.item = p
	return true
}

// IndexOf returns the index of the first occurrence of the
//specified element
func (l *LinkedList) IndexOf(p *collection.LinearObject) int {
	if l.checkNil(p) || !l.checkType(p) {
		return -1
	}
	index := 0
	for node := l.head; node != nil; node = node.next {
		if reflect.DeepEqual(*node.item, *p) {
			return index
		}
		index++
	}
	return -1
}

// Peek returns the head of this queue, or nil if this queue is empty
func (l *LinkedList) First() *collection.LinearObject {
	return l.GetFirst()
}

// Poll returns and removes the head of this queue, or nil if this queue is empty
func (l *LinkedList) Poll() *collection.LinearObject {
	if l.size == 0 {
		return nil
	}
	return l.unlinkFirst()
}

// Add inserts the specified element to the end of this queue.
func (l *LinkedList) Add(p *collection.LinearObject) bool {
	return l.Append(p)
}

// AddFirst inserts the specified element at the front of this deque.
func (l *LinkedList) AddFirst(p *collection.LinearObject) bool {
	if l.checkNil(p) || !l.checkType(p) {
		return false
	}
	return l.linkFirst(p)
}

// AddLast inserts the specified element at the end of this deque.
func (l *LinkedList) AddLast(p *collection.LinearObject) bool {
	if l.checkNil(p) || !l.checkType(p) {
		return false
	}
	return l.linkLast(p)
}

// RemoveFirst removes and returns the head of this deque,
// or returns nil if this deque is empty.
func (l *LinkedList) RemoveFirst() *collection.LinearObject {
	if l.size == 0 {
		return nil
	}
	return l.unlinkFirst()
}

// RemoveLast removes and returns the tail of this deque,
// or returns nil if this deque is empty.
func (l *LinkedList) RemoveLast() *collection.LinearObject {
	if l.size == 0 {
		return nil
	}
	return l.unlinkLast()
}

// GetFirst returns the head of this queue, or nil if this deque is empty.
func (l *LinkedList) GetFirst() *collection.LinearObject {
	if l.size == 0 {
		return nil
	}
	return l.head.item
}

// GetLast returns the tail of this queue, or nil if this deque is empty.
func (l *LinkedList) GetLast() *collection.LinearObject {
	if l.size == 0 {
		return nil
	}
	return l.tail.item
}

// Pop removes and returns the top element of this stack,
// or returns nil if this stack is empty.
func (l *LinkedList) Pop() *collection.LinearObject {
	if l.size == 0 {
		return nil
	}
	return l.unlinkLast()
}

// Push inserts the specified element at the top of this stack.
func (l *LinkedList) Push(p *collection.LinearObject) {
	if l.checkNil(p) || !l.checkType(p) {
		return
	}
	l.linkLast(p)
}

// Peek returns the top element of this stack, or nil if this stack is empty.
func (l *LinkedList) Peek() *collection.LinearObject {
	if l.tail == nil {
		return nil
	}
	return l.tail.item
}

// HashNext returns true if the iteration has more elements.
func (it *LinkedListIterator) HashNext() bool {
	return it.next != nil
}

// Next returns the next element in the iteration.
func (it *LinkedListIterator) Next() *collection.LinearObject {
	if it.HashNext() {
		it.lastRet = it.next
		it.next = it.next.next
		return it.lastRet.item
	}
	return nil
}

func (l *LinkedList) checkNil(p *collection.LinearObject) bool {
	return p == nil || (*p) == nil
}

func (l *LinkedList) checkIndex(index int) bool {
	return index >= 0 && index < l.size
}

func (l *LinkedList) checkType(p *collection.LinearObject) bool {
	return  reflect.TypeOf(*p) == l.t
}

func (l *LinkedList) linkFirst(p *collection.LinearObject) bool {
	node := &Node{p, nil, l.head}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.head.prev = node
		l.head = node
	}
	l.size++
	return true
}

func (l *LinkedList) linkLast(p *collection.LinearObject) bool {
	node := &Node{p, l.tail, nil}
	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
	l.size++
	return true
}

func (l *LinkedList) linkBefore(p *collection.LinearObject, node *Node) bool {
	prev := node.prev
	if prev == nil {
		return l.linkFirst(p)
	} else {
		temp := &Node{p, prev, node}
		prev.next = temp
		node.prev = temp
		l.size++
		return true
	}
}

func (l *LinkedList) unlinkFirst() *collection.LinearObject {
	node := l.head
	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size = 0
	} else {
		l.head = node.next
		l.head.prev = nil
		node.next = nil
		l.size--
	}
	return node.item
}

func (l *LinkedList) unlinkLast() *collection.LinearObject {
	node := l.tail
	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size = 0
	} else {
		l.tail = node.prev
		l.tail.next = nil
		node.prev = nil
		l.size--
	}
	return node.item
}

func (l *LinkedList) unlink(p *Node) bool {
	prev := p.prev
	next := p.next
	if prev == nil {
		l.head = next
	} else {
		prev.next = next
		p.prev = nil
	}
	if next == nil {
		l.tail = prev
	} else {
		next.prev = prev
		p.next = nil
	}
	l.size--
	return true
}

func (l *LinkedList) node(index int) *Node {
	if index < (l.size >> 1) {
		x := l.head
		for i := 0; i < index; i++ {
			x = x.next
		}
		return x
	} else {
		x := l.tail
		for i := l.size - 1; i > index; i-- {
			x = x.prev
		}
		return x
	}
}

func (l *LinkedList) nodeBy(p *collection.LinearObject) *Node {
	for node := l.head; node != nil; node = node.next {
		if reflect.DeepEqual(*node.item, *p) {
			return node
		}
	}
	return nil
}