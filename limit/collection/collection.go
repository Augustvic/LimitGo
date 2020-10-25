package collection

import (
	"reflect"
)

type Linear interface {
	// Size returns the number of elements in this collection.
	Size() int
	// Empty returns true if this collection contains no element.
	Empty() bool
	// GetIterator returns an iterator over the elements in this collection.
	GetIterator() Itr
	// GetType returns type of the elements in this collection.
	GetType() reflect.Type
	// String returns a string representation of this collection.
	String() string
	// Removes all of the elements from this collection.
	Clear() bool
	// Contains returns true if this list contains the specific element.
	Contains(p *Object) bool
}

// List is an interface much similar to java.util.List.
type List interface {
	// Append appends the specified element to the end of this list.
	Append(p *Object) bool
	// Insert the specified element at the specified position in this list.
	Insert(index int, p *Object) bool
	// AddAll appends all of the elements in the specified list to the end of this list.
	AddAll(list *Linear) bool
	// Remove the first occurrence of the specified element from this list.
	Remove(p *Object) bool
	// Removes the element at the specified position in this list.
	RemoveAt(index int) *Object
	// Equals returns true only if the corresponding pairs of the elements
	//in the two lists are equal.
	Equals(list *List) bool
	// Get returns the element at the specified position in this list.
	Get(index int) *Object
	// Set replaces the element at the specified position in this list with
	//the specified element.
	Set(index int, p *Object) bool
	// IndexOf returns the index of the first occurrence of the
	//specified element
	IndexOf(p *Object) int

	// Size returns the number of elements in this collection.
	Size() int
	// Empty returns true if this collection contains no element.
	Empty() bool
	// GetIterator returns an iterator over the elements in this collection.
	GetIterator() Itr
	// GetType returns type of the elements in this collection.
	GetType() reflect.Type
	// String returns a string representation of this collection.
	String() string
	// Removes all of the elements from this collection.
	Clear() bool
	// Contains returns true if this list contains the specific element.
	Contains(p *Object) bool
}

// Queue is an interface much similar to java.util.Queue.
type Queue interface {
	// Peek returns the head of this queue, or nil if this queue is empty
	First() *Object
	// Poll returns and removes the head of this queue, or nil if this queue is empty
	Poll() *Object
	// Add inserts the specified element to the end of this queue.
	Add(p *Object) bool

	// Size returns the number of elements in this collection.
	Size() int
	// Empty returns true if this collection contains no element.
	Empty() bool
	// GetIterator returns an iterator over the elements in this collection.
	GetIterator() Itr
	// GetType returns type of the elements in this collection.
	GetType() reflect.Type
	// String returns a string representation of this collection.
	String() string
	// Removes all of the elements from this collection.
	Clear() bool
	// Contains returns true if this collection contains the specific element.
	Contains(p *Object) bool
}

// Deque is an interface much similar to java.util.Deque.
type Deque interface {
	// AddFirst inserts the specified element at the front of this deque.
	AddFirst(p *Object) bool
	// AddLast inserts the specified element at the end of this deque.
	AddLast(p *Object) bool
	// RemoveFirst removes and returns the head of this deque,
	// or returns nil if this deque is empty.
	RemoveFirst() *Object
	// RemoveLast removes and returns the tail of this deque,
	// or returns nil if this deque is empty.
	RemoveLast() *Object
	// GetFirst returns the head of this queue, or nil if this deque is empty.
	GetFirst() *Object
	// GetLast returns the tail of this queue, or nil if this deque is empty.
	GetLast() *Object

	// Size returns the number of elements in this collection.
	Size() int
	// Empty returns true if this collection contains no element.
	Empty() bool
	// GetIterator returns an iterator over the elements in this collection.
	GetIterator() Itr
	// GetType returns type of the elements in this collection.
	GetType() reflect.Type
	// String returns a string representation of this collection.
	String() string
	// Removes all of the elements from this collection.
	Clear() bool
	// Contains returns true if this collection contains the specific element.
	Contains(p *Object) bool
}

// Stack is an interface much similar to java.util.Stack.
type Stack interface {
	// Pop removes and returns the top element of this stack,
	// or returns nil if this stack is empty.
	Pop() *Object
	// Push inserts the specified element at the top of this stack.
	Push(p *Object)
	// Peek returns the top element of this stack, or nil if this stack is empty.
	Peek() *Object

	// Size returns the number of elements in this collection.
	Size() int
	// Empty returns true if this collection contains no element.
	Empty() bool
	// GetIterator returns an iterator over the elements in this collection.
	GetIterator() Itr
	// GetType returns type of the elements in this collection.
	GetType() reflect.Type
	// String returns a string representation of this collection.
	String() string
	// Removes all of the elements from this collection.
	Clear() bool
	// Contains returns true if this collection contains the specific element.
	Contains(p *Object) bool
}

// Set is an interface much similar to java.util.Set.
type Set interface {
	// Size returns the number of elements in this collection.
	Size() int
	// Empty returns true if this collection contains no element.
	Empty() bool
	// GetIterator returns an iterator over the elements in this collection.
	GetIterator() Itr
	// GetType returns type of the elements in this collection.
	GetType() reflect.Type
	// String returns a string representation of this collection.
	String() string
	// Removes all of the elements from this collection.
	Clear() bool
	// Contains returns true if this list contains the specific element.
	Contains(p *Object) bool

	// Add inserts the specified element to this collection.
	Add(p *Object) bool
	// Remove the first occurrence of the specified element from this collection.
	Remove(p *Object) bool
	// AddAll appends all of the elements in the specified collection to this set.
	AddAll(list *Linear) bool
	// Retains only the elements in this set that are contained in the
	// specified collection.
	RetainAll(list *Linear) bool
	// Removes from this set all of its elements that are contained in the
	// specified collection.
	RemoveAll(list *Linear) bool
	// Equals returns true only if the corresponding pairs of the elements
	//in the two sets are equal.
	Equals(set *Set) bool
}

// An iterator over a linear collection
type Itr interface {
	// HashNext returns true if the iteration has more elements.
	HashNext() bool
	// Next returns the next element in the iteration.
	Next() *Object
	// Remove removes from the underlying collection the last element returned
	// by this iterator.
	Remove() (*Object, bool)
}