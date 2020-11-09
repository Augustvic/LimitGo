package skiplistmap

import (
	"LimitGo/limit/collection"
	"LimitGo/limit/collection/linear/hashset"
	"testing"
)

var mvs *Values
var smvs *Values

func RestartValues() {
	RestartSubMap()
	vs1 := m.Values()
	mvs = (*vs1).(*Values)
	vs2 := sm.Values()
	smvs = (*vs2).(*Values)
}

func TestValuesAll(t *testing.T) {
	TestValues_Clear(t)
	TestValues_Contains(t)
	TestValues_Empty(t)
	TestValues_GetIterator(t)
}

func TestValues_Clear(t *testing.T) {
	RestartValues()
	if mvs.Size() == 0 {
		t.Error("Start operation fail!")
	}
	mvs.Clear()
	if mvs.Size() != 0 || m.Size() != 0 {
		t.Error("Clear operation fail!")
	}
	RestartValues()
	if smvs.Size() == 0 {
		t.Error("Start operation fail!")
	}
	smvs.Clear()
	if smvs.Size() != 0 || sm.Size() != 0 || m.Size() != 1 {
		t.Error("Clear operation fail!")
	}
}

func TestValues_Contains(t *testing.T) {
	RestartValues()
	var st1t collection.Object = Student{100, "st11"}
	var st2t collection.Object = Student{101, "st12"}
	s1t := *hashset.New()
	s1t.Add(&st1t)
	s1t.Add(&st2t)
	var s1to collection.Object = s1t
	if !mvs.Contains(&s1to) {
		t.Error("Contains operation fail!")
	}
	if smvs.Contains(&s1to) {
		t.Error("Contains operation fail!")
	}
}

func TestValues_Empty(t *testing.T) {
	RestartValues()
	if mvs.Empty() {
		t.Error("Start operation fail!")
	}
	mvs.Clear()
	if !mvs.Empty() {
		t.Error("Empty operation fail!")
	}
	RestartValues()
	if smvs.Empty() {
		t.Error("Start operation fail!")
	}
	smvs.Clear()
	if !smvs.Empty() || mvs.Empty() {
		t.Error("Empty operation fail!")
	}
}

func TestValues_GetIterator(t *testing.T) {
	RestartValues()
	it := mvs.GetIterator()
	count := 0
	for it.HashNext() {
		count++
		it.Next()
	}
	if count != 4 {
		t.Error("GetIterator operation fail!")
	}
	it2 := smvs.GetIterator()
	count2 := 0
	for it2.HashNext() {
		count2++
		it2.Next()
	}
	if count2 != 3 {
		t.Error("GetIterator operation fail!")
	}
}