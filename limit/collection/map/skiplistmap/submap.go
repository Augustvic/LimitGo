package skiplistmap

import (
	"LimitGo/limit/collection"
	"reflect"
)

type SubMap struct {
	// underlying map
	sm *SkipListMap
	// lower bound key, or null if from start
	lo *collection.Object
	// upper bound key, or null if to end
	hi *collection.Object
	// inclusion flag for lo
	loInclusive bool
	// inclusion flag for hi
	hiInclusive bool
	// direction
	isDescending bool

	// Lazily initialized view holders
	keySetView   *KeySet
	entrySetView *EntrySet
	valuesView   *Values
}

// Size returns the number of elements in this collection.
func (sm *SubMap) Size() int {
	return 0
}

// Empty returns true if this collection contains no element.
func (sm *SubMap) Empty() bool {
	return false
}

// String returns a string representation of this collection.
func (sm *SubMap) String() string {
	return ""
}

// Removes all of the elements from this collection.
func (sm *SubMap) Clear() bool {
	return false
}

// GetEntryIterator returns iterator of entry.
func (sm *SubMap) GetEntryIterator() collection.EntryItr {
	return nil
}

// ContainsKey returns true if this map contains a mapping for the specified key.
func (sm *SubMap) ContainsKey(key *collection.Object) bool {
	return false
}

// ContainsValue returns true if this map maps one or more keys to the
// specified value.
func (sm *SubMap) ContainsValue(value *collection.Object) bool {
	return false
}

// Get returns the value to which the specified key is mapped, or null
// if this map contains no mapping for the key.
func (sm *SubMap) Get(key *collection.Object) *collection.Object {
	return nil
}

// Put associates the specified value with the specified key, returns old value
// if the specified key has been in this map.
func (sm *SubMap) Put(key *collection.Object, value *collection.Object) (bool, *collection.Object) {
	return true, nil
}

// Remove removes the mapping for a key from this map if it is present.
func (sm *SubMap) Remove(key *collection.Object) *collection.Object {
	return nil
}

// PutAll copies all of the mappings from the specified map to this map.
func (sm *SubMap) PutAll(m *collection.Map) {

}

// KeySet returns a Set view of the keys contained in this map.
func (sm *SubMap) KeySet() *collection.Set {
	return nil
}

// Values returns a List view of the values contained in this map.
func (sm *SubMap) Values() *collection.Linear {
	return nil
}

// EntrySet returns a Set view of the mappings contained in this map.
func (sm *SubMap) EntrySet() *collection.Set {
	return nil
}

// Equals returns true only if the corresponding pairs of the elements
//in the two maps are equal.
func (sm *SubMap) Equals(m *collection.Map) bool {
	return false
}

// SubMap returns a view of the portion of this map whose keys range
// from "fromKey" to "toKey".  If "fromKey" and "toKey" are equal,
// the returned map is empty.)
func (sm *SubMap) SubMap(fromKey *collection.Object, fromInclusive bool, toKey *collection.Object, toInclusive bool) *collection.SortedMap {
	return nil
}

// HeadMap returns a view of the portion of this map whose keys are strictly
// less than toKey.
func (sm *SubMap) HeadMap(toKey *collection.Object, inclusive bool) *collection.SortedMap {
	return nil
}

// TailMap returns a view of the portion of this map whose keys are greater than
// or equal to fromKey.
func (sm *SubMap) TailMap(fromKey *collection.Object, inclusive bool) *collection.SortedMap {
	return nil
}

// SortedKeySet returns a SortedSet view of the keys contained in this map.
func (sm *SubMap) SortedKeySet() *collection.SortedSet {
	return nil
}

// LowerEntry returns a key-value mapping associated with the greatest key
// strictly less than the given key, or nil if there is no such key.
func (sm *SubMap) LowerEntry(key *collection.Object) *collection.Entry {
	return nil
}

// FloorEntry returns a key-value mapping associated with the greatest key
// less than or equal to the given key, or nil if there is no such key.
func (sm *SubMap) FloorEntry(key *collection.Object) *collection.Entry {
	return nil
}

// CeilingEntry returns a key-value mapping associated with the least key
// greater than or equal to the given key, or nil if there is no such key.
func (sm *SubMap) CeilingEntry(key *collection.Object) *collection.Entry {
	return nil
}

// HigherEntry returns a key-value mapping associated with the least key
// strictly greater than the given key, or nil if there is no such key.
func (sm *SubMap) HigherEntry(key *collection.Object) *collection.Entry {
	return nil
}

// Entry returns a key-value mapping associated with the least key
// in this map, or nil if the map is empty.
func (sm *SubMap) FirstEntry() *collection.Entry {
	return nil
}

// LastEntry returns a key-value mapping associated with the greatest
// key in this map, or nil if the map is empty.
func (sm *SubMap) LastEntry() *collection.Entry {
	return nil
}

// PollFirstEntry removes and returns a key-value mapping associated with
// the least key in this map, or nil if the map is empty.
func (sm *SubMap) PollFirstEntry() *collection.Entry {
	return nil
}

// PollLastEntry removes and returns a key-value mapping associated with
// the greatest key in this map, or null if the map is empty.
func (sm *SubMap) PollLastEntry() *collection.Entry {
	return nil
}

func (sm *SubMap) GetKeyType() reflect.Type {
	return (*sm.sm).GetKeyType()
}

func (sm *SubMap) GetValueType() reflect.Type {
	return (*sm.sm).GetValueType()
}
