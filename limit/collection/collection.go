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
	GetIterator() LinearItr
	// GetType returns type of the elements in this collection.
	GetType() reflect.Type
	// String returns a string representation of this collection.
	String() string
	// Removes all of the elements from this collection.
	Clear() bool
}

// List is an interface much similar to java.util.List.
type List interface {
	// Contains returns true if this list contains the specific element.
	Contains(p *LinearObject) bool
	// Append appends the specified element to the end of this list.
	Append(p *LinearObject) bool
	// Insert the specified element at the specified position in this list.
	Insert(index int, p *LinearObject) bool
	// AddAll appends all of the elements in the specified list to the end of this list.
	AddAll(list *Linear) bool
	// Remove the first occurrence of the specified element from this list.
	Remove(p *LinearObject) bool
	// Removes the element at the specified position in this list.
	RemoveAt(index int) *LinearObject
	// Equals returns true only if the corresponding pairs of the elements
	//in the two lists are equal.
	Equals(list *List) bool
	// Get returns the element at the specified position in this list.
	Get(index int) *LinearObject
	// Set replaces the element at the specified position in this list with
	//the specified element.
	Set(index int, p *LinearObject) bool
	// IndexOf returns the index of the first occurrence of the
	//specified element
	IndexOf(p *LinearObject) int

	// Size returns the number of elements in this collection.
	Size() int
	// Empty returns true if this collection contains no element.
	Empty() bool
	// GetIterator returns an iterator over the elements in this collection.
	GetIterator() LinearItr
	// GetType returns type of the elements in this collection.
	GetType() reflect.Type
	// String returns a string representation of this collection.
	String() string
	// Removes all of the elements from this collection.
	Clear() bool
}

// Queue is an interface much similar to java.util.Queue.
type Queue interface {
	// Peek returns the head of this queue, or nil if this queue is empty
	First() *LinearObject
	// Poll returns and removes the head of this queue, or nil if this queue is empty
	Poll() *LinearObject
	// Add inserts the specified element to the end of this queue.
	Add(p *LinearObject) bool

	// Size returns the number of elements in this collection.
	Size() int
	// Empty returns true if this collection contains no element.
	Empty() bool
	// GetIterator returns an iterator over the elements in this collection.
	GetIterator() LinearItr
	// GetType returns type of the elements in this collection.
	GetType() reflect.Type
	// String returns a string representation of this collection.
	String() string
	// Removes all of the elements from this collection.
	Clear() bool
}

// Deque is an interface much similar to java.util.Deque.
type Deque interface {
	// AddFirst inserts the specified element at the front of this deque.
	AddFirst(p *LinearObject) bool
	// AddLast inserts the specified element at the end of this deque.
	AddLast(p *LinearObject) bool
	// RemoveFirst removes and returns the head of this deque,
	// or returns nil if this deque is empty.
	RemoveFirst() *LinearObject
	// RemoveLast removes and returns the tail of this deque,
	// or returns nil if this deque is empty.
	RemoveLast() *LinearObject
	// GetFirst returns the head of this queue, or nil if this deque is empty.
	GetFirst() *LinearObject
	// GetLast returns the tail of this queue, or nil if this deque is empty.
	GetLast() *LinearObject

	// Size returns the number of elements in this collection.
	Size() int
	// Empty returns true if this collection contains no element.
	Empty() bool
	// GetIterator returns an iterator over the elements in this collection.
	GetIterator() LinearItr
	// GetType returns type of the elements in this collection.
	GetType() reflect.Type
	// String returns a string representation of this collection.
	String() string
	// Removes all of the elements from this collection.
	Clear() bool
}

// Stack is an interface much similar to java.util.Stack.
type Stack interface {
	// Pop removes and returns the top element of this stack,
	// or returns nil if this stack is empty.
	Pop() *LinearObject
	// Push inserts the specified element at the top of this stack.
	Push(p *LinearObject)
	// Peek returns the top element of this stack, or nil if this stack is empty.
	Peek() *LinearObject

	// Size returns the number of elements in this collection.
	Size() int
	// Empty returns true if this collection contains no element.
	Empty() bool
	// GetIterator returns an iterator over the elements in this collection.
	GetIterator() LinearItr
	// GetType returns type of the elements in this collection.
	GetType() reflect.Type
	// String returns a string representation of this collection.
	String() string
	// Removes all of the elements from this collection.
	Clear() bool
}

// An iterator over a collection
type LinearItr interface {
	// HashNext returns true if the iteration has more elements.
	HashNext() bool
	// Next returns the next element in the iteration.
	Next() *LinearObject
}