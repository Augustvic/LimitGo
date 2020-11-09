package skiplistmap

import (
	"LimitGo/limit/collection"
	"bytes"
	"encoding/json"
	"fmt"
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
	//// direction
	//isDescending bool

	// Lazily initialized view holders
	keySetView   *KeySet
	entrySetView *EntrySet
	valuesView   *Values
}

type SubMapEntryIterator struct {
	m      *SubMap
	next    *Node
	lastRet *Node
}

// Size returns the number of elements in this collection.
func (m *SubMap) Size() int {
	count := 0
	for n := m.loNode(); m.isBeforeEnd(n); n = n.next {
		count++
	}
	return count
}

// Empty returns true if this collection contains no element.
func (m *SubMap) Empty() bool {
	return m.Size() == 0
}

// String returns a string representation of this collection.
func (m *SubMap) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	it := m.GetEntryIterator()
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
func (m *SubMap) Clear() bool {
	for n := m.loNode(); m.isBeforeEnd(n); n = n.next {
		m.Remove(n.key)
	}
	return true
}

// GetEntryIterator returns iterator of entry.
func (m *SubMap) GetEntryIterator() collection.EntryItr {
	it := SubMapEntryIterator{m, nil, nil}
	it.Init()
	return &it
}

// ContainsKey returns true if this map contains a mapping for the specified key.
func (m *SubMap) ContainsKey(key *collection.Object) bool {
	if m.checkNil(key) {
		return false
	}
	return m.inBounds(key) && m.sm.ContainsKey(key);
}

// ContainsValue returns true if this map maps one or more keys to the
// specified value.
func (m *SubMap) ContainsValue(value *collection.Object) bool {
	if m.checkNil(value) {
		return false
	}
	for n := m.loNode(); m.isBeforeEnd(n); n = n.next {
		v := n.GetValue()
		if v != nil && (*v) != nil && reflect.DeepEqual(*v, *value) {
			return true
		}
	}
	return false
}

// Get returns the value to which the specified key is mapped, or null
// if this map contains no mapping for the key.
func (m *SubMap) Get(key *collection.Object) *collection.Object {
	if m.checkNil(key) || !m.inBounds(key) {
		return nil
	}
	return m.sm.Get(key)
}

// Put associates the specified value with the specified key, returns old value
// if the specified key has been in this map.
func (m *SubMap) Put(key *collection.Object, value *collection.Object) (bool, *collection.Object) {
	if m.checkNil(key) {
		return false, nil
	}
	if m.checkNil(value) {
		return false, nil
	}
	return m.sm.Put(key, value)
}

// Remove removes the mapping for a key from this map if it is present.
func (m *SubMap) Remove(key *collection.Object) *collection.Object {
	if m.checkNil(key) || !m.inBounds(key) {
		return nil
	}
	return m.sm.Remove(key)
}

// PutAll copies all of the mappings from the specified map to this map.
func (m *SubMap) PutAll(m2 *collection.Map) {
	if m2 == nil || (*m2) == nil || (*m2).Size() == 0 {
		return
	}
	it := (*m2).GetEntryIterator()
	for it.HashNext() {
		entry := (*it.Next()).(collection.Entry)
		m.Put(entry.GetKey(), entry.GetValue())
	}
}

// KeySet returns a Set view of the keys contained in this map.
func (m *SubMap) KeySet() *collection.Set {
	if m.keySetView != nil {
		var s collection.Set = m.keySetView
		return &s
	}
	var t collection.SortedMap = m
	m.keySetView = &KeySet{&t}
	var s collection.Set= m.keySetView
	return &s
}

// Values returns a List view of the values contained in this map.
func (m *SubMap) Values() *collection.Linear {
	if m.valuesView != nil {
		var s collection.Linear = m.valuesView
		return &s
	}
	var t collection.SortedMap = m
	m.valuesView = &Values{&t}
	var s collection.Linear = m.valuesView
	return &s
}

// EntrySet returns a Set view of the mappings contained in this map.
func (m *SubMap) EntrySet() *collection.Set {
	if m.entrySetView != nil {
		var s collection.Set = m.entrySetView
		return &s
	}
	var t collection.SortedMap = m
	m.entrySetView = &EntrySet{&t}
	var s collection.Set = m.entrySetView
	return &s
}

// Equals returns true only if the corresponding pairs of the elements
//in the two maps are equal.
func (m *SubMap) Equals(m2 *collection.Map) bool {
	if m2 == nil || (*m2) == nil || m.Size() != (*m2).Size() {
		return false
	}
	it := (*m).GetEntryIterator()
	for it.HashNext() {
		entry := (*it.Next()).(collection.Entry)
		if !reflect.DeepEqual(*m.Get(entry.GetKey()), *entry.GetValue()) {
			return false
		}
	}
	return true
}

// SubMap returns a view of the portion of this map whose keys range
// from "fromKey" to "toKey".  If "fromKey" and "toKey" are equal,
// the returned map is empty.)
func (m *SubMap) SubMap(fromKey *collection.Object, fromInclusive bool, toKey *collection.Object, toInclusive bool) *collection.SortedMap {
	if !m.checkSubMap(fromKey, fromInclusive, toKey, toInclusive) {
		return nil
	}
	if m.lo != nil {
		if fromKey == nil {
			fromKey = m.lo
			fromInclusive = m.loInclusive
		} else {
			if m.sm.precede(fromKey, m.lo) || (reflect.DeepEqual(*fromKey, *m.lo) && !m.loInclusive && fromInclusive) {
				return nil
			}
		}
	}
	if m.hi != nil {
		if toKey == nil {
			toKey = m.lo
			toInclusive = m.hiInclusive
		} else {
			if m.sm.precede(m.hi, toKey) || (reflect.DeepEqual(*toKey, *m.hi) && !m.hiInclusive && toInclusive) {
				return nil
			}
		}
	}
	p := &SubMap{m.sm, fromKey, toKey, fromInclusive, toInclusive, nil, nil, nil}
	var ret collection.SortedMap = p
	return &ret
}

// HeadMap returns a view of the portion of this map whose keys are strictly
// less than toKey.
func (m *SubMap) HeadMap(toKey *collection.Object, inclusive bool) *collection.SortedMap {
	if toKey != nil && !m.inBounds(toKey) {
		return nil
	}
	return m.SubMap(nil, false, toKey, inclusive)
}

// TailMap returns a view of the portion of this map whose keys are greater than
// or equal to fromKey.
func (m *SubMap) TailMap(fromKey *collection.Object, inclusive bool) *collection.SortedMap {
	if fromKey != nil && !m.inBounds(fromKey) {
		return nil
	}
	return m.SubMap(fromKey, inclusive, nil, false)
}

// SortedKeySet returns a SortedSet view of the keys contained in this map.
func (m *SubMap) SortedKeySet() *collection.SortedSet {
	ss := (*m.KeySet()).(collection.SortedSet)
	return &ss
}

// LowerEntry returns a key-value mapping associated with the greatest key
// strictly less than the given key, or nil if there is no such key.
func (m *SubMap) LowerEntry(key *collection.Object) *collection.Entry {
	if m.checkNil(key) {
		return nil
	}
	return m.getNearEntry(key, LT)
}

// FloorEntry returns a key-value mapping associated with the greatest key
// less than or equal to the given key, or nil if there is no such key.
func (m *SubMap) FloorEntry(key *collection.Object) *collection.Entry {
	if m.checkNil(key) {
		return nil
	}
	return m.getNearEntry(key, LT|EQ)
}

// CeilingEntry returns a key-value mapping associated with the least key
// greater than or equal to the given key, or nil if there is no such key.
func (m *SubMap) CeilingEntry(key *collection.Object) *collection.Entry {
	if m.checkNil(key) {
		return nil
	}
	return m.getNearEntry(key, GT|EQ)
}

// HigherEntry returns a key-value mapping associated with the least key
// strictly greater than the given key, or nil if there is no such key.
func (m *SubMap) HigherEntry(key *collection.Object) *collection.Entry {
	if m.checkNil(key) {
		return nil
	}
	return m.getNearEntry(key, GT)
}

// Entry returns a key-value mapping associated with the least key
// in this map, or nil if the map is empty.
func (m *SubMap) FirstEntry() *collection.Entry {
	var t collection.Entry = m.loNode()
	if t == nil {
		return nil
	} else {
		return &t
	}
}

// LastEntry returns a key-value mapping associated with the greatest
// key in this map, or nil if the map is empty.
func (m *SubMap) LastEntry() *collection.Entry {
	var t collection.Entry = m.hiNode()
	if t == nil {
		return nil
	} else {
		return &t
	}
}

// PollFirstEntry removes and returns a key-value mapping associated with
// the least key in this map, or nil if the map is empty.
func (m *SubMap) PollFirstEntry() *collection.Entry {
	return m.removeLowest()
}

// PollLastEntry removes and returns a key-value mapping associated with
// the greatest key in this map, or null if the map is empty.
func (m *SubMap) PollLastEntry() *collection.Entry {
	return m.removeHighest()
}

func (m *SubMap) loNode() *Node {
	if m.lo == nil {
		return m.sm.findFirst()
	} else if m.loInclusive {
		return m.sm.findNear(m.lo, GT|EQ)
	} else {
		return m.sm.findNear(m.lo, GT)
	}
}

func (m *SubMap) hiNode() *Node {
	if m.hi == nil {
		return m.sm.findLast()
	} else if m.hiInclusive {
		return m.sm.findNear(m.hi, LT|EQ)
	} else {
		return m.sm.findNear(m.hi, LT)
	}
}

func (m *SubMap) isBeforeEnd(node *Node) bool {
	if node == nil {
		return false
	}
	if m.hi == nil {
		return true
	}
	k := node.key
	if m.sm.precede(m.hi, k) || (reflect.DeepEqual(*k, *m.hi) && !m.hiInclusive) {
		return false
	}
	return true
}

func (m *SubMap) inBounds(p *collection.Object) bool {
	if m.tooLow(p) || m.tooHigh(p) {
		return false
	}
	return true
}

func (m *SubMap) tooLow(p *collection.Object) bool {
	if m.lo != nil && ((reflect.DeepEqual(*p, *m.lo) && !m.loInclusive) || (m.sm.precede(p, m.lo))) {
		return true
	}
	return false
}

func (m *SubMap) tooHigh(p *collection.Object) bool {
	if m.hi != nil && ((reflect.DeepEqual(*p, *m.hi) && !m.hiInclusive) || (m.sm.precede(m.hi, p))) {
		return true
	}
	return false
}

func (m *SubMap) getNearEntry(key *collection.Object, rel int) *collection.Entry {
	if m.tooLow(key) {
		if (rel & LT) != 0 {
			return nil
		}
		return m.lowestEntry()
	}
	if m.tooHigh(key) {
		if (rel & LT) != 0 {
			return m.highestEntry()
		}
		return nil
	}
	n := m.sm.findNear(key, rel)
	if n == nil || !m.inBounds(n.key) {
		return nil
	}
	var e collection.Entry = n
	return &e
}

func (m *SubMap) removeHighest() *collection.Entry {
	var n collection.Entry = m.hiNode()
	if n == nil {
		return nil
	} else {
		m.Remove(n.GetKey())
		return &n
	}
}

func (m *SubMap) removeLowest() *collection.Entry {
	var n collection.Entry = m.loNode()
	if n == nil {
		return nil
	} else {
		m.Remove(n.GetKey())
		return &n
	}
}

func (m *SubMap) lowestEntry() *collection.Entry {
	n := m.loNode()
	if !m.isBeforeEnd(n) {
		return nil
	}
	var e collection.Entry = n
	return &e
}

func (m *SubMap) highestEntry() *collection.Entry {
	n := m.hiNode()
	if n == nil || !m.inBounds(n.key) {
		return nil
	}
	var e collection.Entry = n
	return &e
}

func (m *SubMap) checkNil(p *collection.Object) bool {
	return p == nil || (*p) == nil
}

func (m *SubMap) checkSubMap(fromKey *collection.Object, fromInclusive bool, toKey *collection.Object, toInclusive bool) bool {
	if fromKey != nil && !m.inBounds(fromKey) {
		return false
	}
	if toKey != nil && !m.inBounds(toKey) {
		return false
	}
	return true
}

// HashNext returns true if the iteration has more elements.
func (it * SubMapEntryIterator) HashNext() bool {
	return it.next != nil
}

// Next returns the next element in the iteration.
func (it * SubMapEntryIterator) Next() *collection.Entry {
	if it.HashNext() {
		it.lastRet = it.next
		it.next = it.next.next
		if it.next != nil && it.m.tooHigh(it.next.key) {
			it.next = nil
		}
		var t collection.Entry =  it.lastRet
		return &t
	}
	return nil
}

// Remove removes from the underlying collection the last element returned
// by this iterator.
func (it * SubMapEntryIterator) Remove() (*collection.Entry, bool) {
	if it.lastRet == nil {
		return nil, false
	}
	var last collection.Entry = it.lastRet
	it.m.sm.doRemove(last.GetKey())
	it.lastRet = nil
	return &last, true
}

func (it * SubMapEntryIterator) Init() {
	next := it.m.loNode()
	if next == nil || !it.m.inBounds(next.key) {
		it.next = nil
	} else {
		it.next = next
	}
}