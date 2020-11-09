package skiplistmap

import (
	"LimitGo/limit/collection"
	"LimitGo/limit/collection/linear/hashset"
	"reflect"
	"testing"
)

var mes *EntrySet
var smes *EntrySet

func RestartEntrySet() {
	RestartSubMap()
	es1 := m.EntrySet()
	mes = (*es1).(*EntrySet)
	es2 := sm.EntrySet()
	smes = (*es2).(*EntrySet)
}

func TestEntrySetAll(t *testing.T) {
	TestEntrySet_Clear(t)
	TestEntrySet_Contains(t)
	TestEntrySet_Empty(t)
	TestEntrySet_Equals(t)
	TestEntrySet_GetIterator(t)
	TestEntrySet_Remove(t)
	TestEntrySet_RemoveAll(t)
	TestEntrySet_RetainAll(t)
}

func TestEntrySet_Clear(t *testing.T) {
	RestartEntrySet()
	if mes.Size() == 0 {
		t.Error("Start operation fail!")
	}
	mes.Clear()
	if mes.Size() != 0 || m.Size() != 0 {
		t.Error("Clear operation fail!")
	}
	RestartEntrySet()
	if smes.Size() == 0 {
		t.Error("Start operation fail!")
	}
	smes.Clear()
	if smes.Size() != 0 || sm.Size() != 0 || m.Size() != 1 {
		t.Error("Clear operation fail!")
	}
}

func TestEntrySet_Contains(t *testing.T) {
	RestartEntrySet()
	var tt2 collection.Object = Teacher{2, "t2", 0}
	var tst11 collection.Object = Student{100, "st11"}
	var tst12 collection.Object = Student{101, "st12"}
	ts1 := *hashset.New()
	ts1.Add(&tst11)
	ts1.Add(&tst12)
	var s1to collection.Object = ts1
	var entry collection.Object = &Node{&tt2, &s1to, nil}
	if !mes.Contains(&entry) {
		t.Error("Contains operation fail!")
	}
	if smes.Contains(&entry) {
		t.Error("Contains operation fail!")
	}
}

func TestEntrySet_Empty(t *testing.T) {
	RestartEntrySet()
	if mes.Empty() {
		t.Error("Start operation fail!")
	}
	mes.Clear()
	if !mes.Empty() {
		t.Error("Empty operation fail!")
	}
	RestartEntrySet()
	if smes.Empty() {
		t.Error("Start operation fail!")
	}
	smes.Clear()
	if !smes.Empty() || mes.Empty() {
		t.Error("Empty operation fail!")
	}
}

func TestEntrySet_Equals(t *testing.T) {
	RestartEntrySet()
	var tt2 collection.Object
	var tt4 collection.Object
	var tt6 collection.Object
	var tt8 collection.Object
	var tst11 collection.Object
	var tst12 collection.Object
	var tst2 collection.Object
	var tst3 collection.Object
	var tst4 collection.Object
	var ts1 collection.Object
	var ts2 collection.Object
	var ts3 collection.Object
	var ts4 collection.Object
	var tkt reflect.Type
	var tvt reflect.Type
	var tm *SkipListMap
	tt2 = Teacher{2, "t2", 0}
	tt4 = Teacher{4, "t4", 0}
	tt6 = Teacher{6, "t6", 0}
	tt8 = Teacher{8, "t8", 0}
	tst11 = Student{100, "st11"}
	tst12 = Student{101, "st12"}
	tst2 = Student{102, "st2"}
	tst3 = Student{103, "st3"}
	tst4 = Student{104, "st4"}
	ts1 = *hashset.New()
	ts2 = *hashset.New()
	ts3 = *hashset.New()
	ts4 = *hashset.New()
	ts1s := ts1.(hashset.HashSet)
	ts2s := ts2.(hashset.HashSet)
	ts3s := ts3.(hashset.HashSet)
	ts4s := ts4.(hashset.HashSet)
	ts1s.Add(&tst11)
	ts1s.Add(&tst12)
	ts2s.Add(&tst2)
	ts3s.Add(&tst3)
	ts4s.Add(&tst4)
	tkt = reflect.TypeOf(tt2)
	tvt = reflect.TypeOf(ts1)
	tm = New(tkt, tvt, precede)
	tm.Put(&tt2, &ts1)
	tm.Put(&tt4, &ts2)
	tm.Put(&tt6, &ts3)
	tm.Put(&tt8, &ts4)
	if !mes.Equals(tm.EntrySet()) {
		t.Error("Equals operation fail!")
	}
}

func TestEntrySet_GetIterator(t *testing.T) {
	RestartEntrySet()
	it := mes.GetIterator()
	count := 0
	for it.HashNext() {
		count++
		it.Next()
	}
	if count != 4 {
		t.Error("GetIterator operation fail!")
	}
	it2 := smes.GetIterator()
	count2 := 0
	for it2.HashNext() {
		count2++
		it2.Next()
	}
	if count2 != 3 {
		t.Error("GetIterator operation fail!")
	}
}

func TestEntrySet_Remove(t *testing.T) {
	RestartEntrySet()
	var tt2 collection.Object = Teacher{2, "t2", 0}
	var tst11 collection.Object = Student{100, "st11"}
	var tst12 collection.Object = Student{101, "st12"}
	ts1 := *hashset.New()
	ts1.Add(&tst11)
	ts1.Add(&tst12)
	var s1to collection.Object = ts1
	var entry collection.Object = &Node{&tt2, &s1to, nil}
	mes.Remove(&entry)
	if mes.Contains(&entry) || mes.Size() != 3 {
		t.Error("Remove operation fail!")
	}
}

func TestEntrySet_RemoveAll(t *testing.T) {
	RestartEntrySet()
	var tt2 collection.Object = Teacher{2, "t2", 0}
	var tst11 collection.Object = Student{100, "st11"}
	var tst12 collection.Object = Student{101, "st12"}
	ts1 := *hashset.New()
	ts1.Add(&tst11)
	ts1.Add(&tst12)
	var s1to collection.Object = ts1
	var entry collection.Object = &Node{&tt2, &s1to, nil}
	var s collection.Linear = hashset.New()
	s.Add(&entry)
	mes.RemoveAll(&s)
	if mes.Contains(&entry) || mes.Size() != 3 {
		t.Error("RemoveAll operation fail!")
	}
}

func TestEntrySet_RetainAll(t *testing.T) {
	RestartEntrySet()
	var s collection.Linear = hashset.New()
	mes.RetainAll(&s)
	if mes.Size() != 0 || smes.Size() != 0 {
		t.Error("RetainAll operation fail!")
	}
}