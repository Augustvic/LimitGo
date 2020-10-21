package collections

import (
	"reflect"
)

// An iterator over a list
type ListItr interface {

	// HashNext returns true if the iteration has more elements.
	HashNext() bool

	// Next returns the next element in the iteration.
	Next() *ListObject
}

// List is a interface much similar to java.util.List.
type List interface {

	// Size returns the number of elements in this list.
	Size() int

	// Empty returns true if this list contains no element.
	Empty() bool

	// Contains returns true if this list contains the specific element.
	Contains(p *ListObject) bool

	// Append appends the specified element to the end of this list.
	Append(p *ListObject)

	// Insert the specified element at the specified position in this list.
	Insert(index int, p *ListObject)

	// Appends all of the elements in the specified list to the end of this list.
	AddAll(list List)

	// Remove the first occurrence of the specified element from this list.
	Remove(p *ListObject)

	// Removes the element at the specified position in this list.
	RemoveAt(index int)

	// Removes all of the elements from this list.
	Clear()

	// Equals returns true only if the corresponding pairs of the elements
	//in the two lists are equal.
	Equals(list List) bool

	// String returns a string representation of this list.
	String() string

	// Get returns the element at the specified position in this list.
	Get(index int) *ListObject

	// Set replaces the element at the specified position in this list with
	//the specified element.
	Set(index int, p *ListObject)

	// IndexOf returns the index of the first occurrence of the
	//specified element
	IndexOf(p *ListObject) int

	// GetIterator returns an iterator over the elements in this list.
	GetIterator() ListItr

	// GetType returns type of the elements in this list.
	GetType() reflect.Type
}