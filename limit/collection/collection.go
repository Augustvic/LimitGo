package collection

import (
	"reflect"
)

type Collection interface {
	// Size returns the number of elements in this collection.
	Size() int
	// Empty returns true if this collection contains no element.
	Empty() bool
	// String returns a string representation of this collection.
	String() string
	// Removes all of the elements from this collection.
	Clear() bool
}

// Interface of linear data structure.
type Linear interface {
	// GetIterator returns an iterator over the elements in this collection.
	GetIterator() Itr
	// Contains returns true if this list contains the specific element.
	Contains(p *Object) bool
	// GetType returns type of the elements in this collection.
	GetType() reflect.Type
	// All Methods from Collection
	Collection
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
	// All Methods from Linear
	Linear
}

// Queue is an interface much similar to java.util.Queue.
type Queue interface {
	// Peek returns the head of this queue, or nil if this queue is empty
	First() *Object
	// Poll returns and removes the head of this queue, or nil if this queue is empty
	Poll() *Object
	// Add inserts the specified element to the end of this queue.
	Add(p *Object) bool
	// All Methods from Linear
	Linear
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
	// All Methods from Linear
	Linear
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
	// All Methods from Linear
	Linear
}

// Set is an interface much similar to java.util.Set.
type Set interface {
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
	// All Methods from Linear
	Linear
}

type SortedSet interface {
	// SubSet returns a view of the portion of this set whose elements range from
	// fromElement to toElement.  If fromElement and toElement are equal,
	// the returned set is empty unless fromInclusive and toInclusive are both true.
	SubSet(fromElement *Object, fromInclusive bool, toElement *Object, toInclusive bool) *SortedSet
	// HeadSet returns a view of the portion of this set whose elements are
	// less than (or equal to, if inclusive is true) toElement
	HeadSet(toElement *Object, inclusive bool) *SortedSet
	// TailSet returns a view of the portion of this set whose elements are
	// greater than (or equal to, if inclusive is true) fromElement.
	TailSet(fromElement *Object, inclusive bool) *SortedSet
	// First returns the first (lowest) element currently in this set.
	First() *Object
	// Last returns the last (highest) element currently in this set.
	Last() *Object
	// Lower returns the greatest element in this set strictly less than the given
	// element, or null if there is no such element.
	Lower(p *Object) *Object
	// Floor returns the greatest element in this set less than or equal to
	// the given element, or null if there is no such element.
	Floor(p *Object) *Object
	// Ceiling returns the least element in this set greater than or equal to
	// the given element, or null if there is no such element.
	Ceiling(p *Object) *Object
	// Higher returns the least element in this set strictly greater than the
	// given element, or null if there is no such element.
	Higher(p *Object) *Object
	// PollFirst retrieves and removes the first (lowest) element, or
	// returns null if this set is empty.
	PollFirst() *Object
	// PollLast retrieves and removes the last (highest) element, or
	// returns null if this set is empty.
	PollLast() *Object
	// DescendingSet returns a reverse order view of the elements contained
	// in this set.
	DescendingSet() *SortedSet
	// All Methods from Set
	Set
}

// Map is an interface much similar to java.util.Map.
type Map interface {
	// GetEntryIterator returns iterator of entry.
	GetEntryIterator() EntryItr
	// ContainsKey returns true if this map contains a mapping for the specified key.
	ContainsKey(key *Object) bool
	// ContainsValue returns true if this map maps one or more keys to the
	// specified value.
	ContainsValue(value *Object) bool
	// Get returns the value to which the specified key is mapped, or null
	// if this map contains no mapping for the key.
	Get(key *Object) *Object
	// Associates the specified value with the specified key in this map.
	Put(key *Object, value *Object) (bool, *Object)
    // Remove removes the mapping for a key from this map if it is present.
	Remove(key *Object) *Object
	// PutAll copies all of the mappings from the specified map to this map.
	PutAll(m *Map)
	// KeySet returns a Set view of the keys contained in this map.
	KeySet() *Set
	// Values returns a List view of the values contained in this map.
	Values() *Linear
	// EntrySet returns a Set view of the mappings contained in this map.
	EntrySet() *Set
	// Equals returns true only if the corresponding pairs of the elements
	//in the two maps are equal.
	Equals(list *Map) bool
	// All Methods from Collection
	Collection
}

type SortedMap interface {
	// SubMap returns a view of the portion of this map whose keys range
	// from "fromKey" to "toKey".  If "fromKey" and "toKey" are equal,
	// the returned map is empty.)
	SubMap(fromKey *Object, fromInclusive bool, toKey *Object, toInclusive bool) *SortedMap
	// HeadMap returns a view of the portion of this map whose keys are strictly
	// less than toKey.
	HeadMap(toKey *Object, inclusive bool) *SortedMap
	// TailMap returns a view of the portion of this map whose keys are greater than
	// or equal to fromKey.
	TailMap(fromKey *Object, inclusive bool) *SortedMap
	// SortedKeySet returns a SortedSet view of the keys contained in this map.
	SortedKeySet() *SortedSet
	// LowerEntry returns a key-value mapping associated with the greatest key
	// strictly less than the given key, or nil if there is no such key.
	LowerEntry(key *Object) *Entry
	// FloorEntry returns a key-value mapping associated with the greatest key
	// less than or equal to the given key, or nil if there is no such key.
	FloorEntry(key *Object) *Entry
	// CeilingEntry returns a key-value mapping associated with the least key
	// greater than or equal to the given key, or nil if there is no such key.
	CeilingEntry(key *Object) *Entry
	// HigherEntry returns a key-value mapping associated with the least key
	// strictly greater than the given key, or nil if there is no such key.
	HigherEntry(key *Object) *Entry
	// Entry returns a key-value mapping associated with the least key
	// in this map, or nil if the map is empty.
	FirstEntry() *Entry
	// LastEntry returns a key-value mapping associated with the greatest
	// key in this map, or nil if the map is empty.
	LastEntry() *Entry
	// PollFirstEntry removes and returns a key-value mapping associated with 
	// the least key in this map, or nil if the map is empty. 
	PollFirstEntry() *Entry
	// PollLastEntry removes and returns a key-value mapping associated with
	// the greatest key in this map, or null if the map is empty.
	PollLastEntry() *Entry
	// All Methods from Map
	Map
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

// An iterator over an map
type EntryItr interface {
	// HashNext returns true if the iteration has more elements.
	HashNext() bool
	// Next returns the next element in the iteration.
	Next() *Entry
	// Remove removes from the underlying collection the last element returned
	// by this iterator.
	Remove() (*Entry, bool)
}