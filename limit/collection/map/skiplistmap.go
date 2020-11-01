package _map

import (
	"LimitGo/limit/collection"
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

// Node hold keys and values, and are singly linked in sorted order.
type Node struct {
	key   *collection.Object
	value *collection.Object
	next  *Node
}

// Index nodes represent the levels of the skip list.
type Index struct {
	node  *Node
	down  *Index
	right *Index
}

// HeadIndex heading each level keep track of their level.
type HeadIndex struct {
	level int
	Index
}

type KeySet struct {
	sm *SkipListMap
}

type EntrySet struct {
	sm *SkipListMap
}

type Values struct {
	sm *SkipListMap
}

// Combination of skip list map and hash map
type SkipListMap struct {
	m       map[collection.Object]*collection.Object
	head    *HeadIndex
	precede func(p1 *collection.Object, p2 *collection.Object) bool
	kt      reflect.Type
	vt      reflect.Type

	keySet   *KeySet
	entrySet *EntrySet
	values   *Values
}

type EntryIterator struct {
	sm      *SkipListMap
	next    *collection.Entry
	lastRet *collection.Entry
}

func New(kt reflect.Type, vt reflect.Type, precede func(p1 *collection.Object, p2 *collection.Object) bool) *SkipListMap {
	m := make(map[collection.Object]*collection.Object)
	head := &HeadIndex{1, Index{nil, nil, nil}}
	return &SkipListMap{m, head, precede, kt, vt, nil, nil, nil}
}

// Size returns the number of elements in this collection.
func (sm *SkipListMap) Size() int {
	return len(sm.m)
}

// Empty returns true if this collection contains no element.
func (sm *SkipListMap) Empty() bool {
	return len(sm.m) == 0
}

// String returns a string representation of this collection.
func (sm *SkipListMap) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	it := sm.GetEntryIterator()
	for it.HashNext() {
		entry := (*it.Next()).(collection.Entry)
		if buf.Len() > len("{") {
			buf.WriteByte(',')
		}
		key := entry.GetKey()
		value := entry.GetValue()
		var s string
		k, err1 := json.Marshal(*key)
		v, err2 := json.Marshal(*value)
		if err1 == nil && err2 == nil {
			s = string(k) + "=" + string(v)
		} else {
			s = "nil"
		}
		_, _ = fmt.Fprint(&buf, s)
	}
	buf.WriteByte('}')
	return buf.String()
}

// Removes all of the elements from this collection.
func (sm *SkipListMap) Clear() bool {
	sm.m = make(map[collection.Object]*collection.Object)
	sm.head = &HeadIndex{1, Index{nil, nil, nil}}
	sm.keySet = &KeySet{sm}
	sm.entrySet = &EntrySet{sm}
	sm.values = &Values{sm}
	return true
}

// GetEntryIterator returns iterator of entry.
func (sm *SkipListMap) GetEntryIterator() collection.Itr {
	return nil
}

// ContainsKey returns true if this map contains a mapping for the specified key.
func (sm *SkipListMap) ContainsKey(key *collection.Object) bool {
	if sm.checkNil(key) || !sm.checkKeyType(key) {
		return false
	}
	return sm.doGet(key) != nil
}

// ContainsValue returns true if this map maps one or more keys to the
// specified value.
func (sm *SkipListMap) ContainsValue(value *collection.Object) bool {
	if sm.checkNil(value) || !sm.checkValueType(value) {
		return false
	}
	for node := sm.findFirst(); node != nil; node = node.next {
		p := node.value
		if reflect.DeepEqual(*value, *p) {
			return true
		}
	}
	return false
}

// Get returns the value to which the specified key is mapped, or null
// if this map contains no mapping for the key.
func (sm *SkipListMap) Get(key *collection.Object) *collection.Object {
	if sm.checkNil(key) || !sm.checkKeyType(key) {
		return nil
	}
	return sm.doGet(key)
}

// Put associates the specified value with the specified key, returns old value
// if the specified key has been in this map.
func (sm *SkipListMap) Put(key *collection.Object, value *collection.Object) (bool, *collection.Object) {
	if sm.checkNil(key) || !sm.checkKeyType(key) {
		return false, nil
	}
	if sm.checkNil(value) || !sm.checkValueType(value) {
		return false, nil
	}
	return true, sm.doPut(key, value)
}

// Remove removes the mapping for a key from this map if it is present.
func (sm *SkipListMap) Remove(key *collection.Object) *collection.Object {
	if sm.checkNil(key) || !sm.checkKeyType(key) {
		return nil
	}
	return sm.doRemove(key)
}

// PutAll copies all of the mappings from the specified map to this map.
func (sm *SkipListMap) PutAll(m *collection.Map) {
	if m == nil || (*m) == nil || (*m).Size() == 0 {
		return
	}
	it := (*m).GetEntryIterator()
	for it.HashNext() {
		entry := (*it.Next()).(collection.Entry)
		sm.Put(entry.GetKey(), entry.GetValue())
	}
}

// KeySet returns a Set view of the keys contained in this map.
func (sm *SkipListMap) KeySet() *collection.Set {
	return nil
}

// Values returns a List view of the values contained in this map.
func (sm *SkipListMap) Values() *collection.List {
	return nil
}

// EntrySet returns a Set view of the mappings contained in this map.
func (sm *SkipListMap) EntrySet() *collection.Set {
	return nil
}

// Equals returns true only if the corresponding pairs of the elements
//in the two maps are equal.
func (sm *SkipListMap) Equals(m *collection.Map) bool {
	if m == nil || (*m) == nil || (*m).Size() == 0 || sm.Size() != (*m).Size() {
		return false
	}
	it := (*m).GetEntryIterator()
	for it.HashNext() {
		entry := (*it.Next()).(collection.Entry)
		if !reflect.DeepEqual(*sm.Get(entry.GetKey()), *entry.GetValue()) {
			return false
		}
	}
	return true
}

// SubMap returns a view of the portion of this map whose keys range
// from "fromKey" to "toKey".  If "fromKey" and "toKey" are equal,
// the returned map is empty.)
func (sm *SkipListMap) SubMap(fromKey *collection.Object, fromInclusive bool, toKey *collection.Object, toInclusive bool) *collection.SortedMap {
	var t collection.SortedMap = &SubMap{sm, fromKey, toKey, fromInclusive, toInclusive, false, nil, nil, nil}
	return &t
}

// HeadMap returns a view of the portion of this map whose keys are strictly
// less than toKey.
func (sm *SkipListMap) HeadMap(toKey *collection.Object, inclusive bool) *collection.SortedMap {
	if sm.checkNil(toKey) || !sm.checkKeyType(toKey) {
		return nil
	}
	var t collection.SortedMap = &SubMap{sm, nil, toKey, false, inclusive, false, nil, nil, nil}
	return &t
}

// TailMap returns a view of the portion of this map whose keys are greater than
// or equal to fromKey.
func (sm *SkipListMap) TailMap(fromKey *collection.Object, inclusive bool) *collection.SortedMap {
	if sm.checkNil(fromKey) || !sm.checkKeyType(fromKey) {
		return nil
	}
	var t collection.SortedMap = &SubMap{sm, fromKey, nil, inclusive, false, false, nil, nil, nil}
	return &t
}

// SortedKeySet returns a SortedSet view of the keys contained in this map.
func (sm *SkipListMap) SortedKeySet() *collection.SortedSet {
	return nil
}

// LowerEntry returns a key-value mapping associated with the greatest key
// strictly less than the given key, or nil if there is no such key.
func (sm *SkipListMap) LowerEntry(key *collection.Object) *collection.Entry {
	return nil
}

// FloorEntry returns a key-value mapping associated with the greatest key
// less than or equal to the given key, or nil if there is no such key.
func (sm *SkipListMap) FloorEntry(key *collection.Object) *collection.Entry {
	return nil
}

// CeilingEntry returns a key-value mapping associated with the least key
// greater than or equal to the given key, or nil if there is no such key.
func (sm *SkipListMap) CeilingEntry(key *collection.Object) *collection.Entry {
	return nil
}

// HigherEntry returns a key-value mapping associated with the least key
// strictly greater than the given key, or nil if there is no such key.
func (sm *SkipListMap) HigherEntry(key *collection.Object) *collection.Entry {
	return nil
}

// Entry returns a key-value mapping associated with the least key
// in this map, or nil if the map is empty.
func (sm *SkipListMap) FirstEntry() *collection.Entry {
	return nil
}

// LastEntry returns a key-value mapping associated with the greatest
// key in this map, or nil if the map is empty.
func (sm *SkipListMap) LastEntry() *collection.Entry {
	return nil
}

// PollFirstEntry removes and returns a key-value mapping associated with
// the least key in this map, or nil if the map is empty.
func (sm *SkipListMap) PollFirstEntry() *collection.Entry {
	return nil
}

// PollLastEntry removes and returns a key-value mapping associated with
// the greatest key in this map, or null if the map is empty.
func (sm *SkipListMap) PollLastEntry() *collection.Entry {
	return nil
}

func (sm *SkipListMap) checkNil(p *collection.Object) bool {
	return p == nil || (*p) == nil
}

func (sm *SkipListMap) checkKeyType(p *collection.Object) bool {
	return reflect.TypeOf(*p) == sm.kt
}

func (sm *SkipListMap) checkValueType(p *collection.Object) bool {
	return reflect.TypeOf(*p) == sm.vt
}

func (sm *SkipListMap) doGet(key *collection.Object) *collection.Object {
	return nil
}

func (sm *SkipListMap) doPut(key *collection.Object, value *collection.Object) *collection.Object{
	return nil
}

func (sm *SkipListMap) doRemove(key *collection.Object) *collection.Object {
	return nil
}

func (sm *SkipListMap) findFirst() *Node {
	return nil
}


