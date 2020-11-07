package skiplistmap

import (
	"LimitGo/limit/collection"
	"LimitGo/limit/collection/linear/set"
	"reflect"
	"strconv"
	"testing"
)

type Student struct {
	Id int
	Name string
}

type Teacher struct {
	Id int
	Name string
	Sex int
}

var precede = func(p1 *collection.Object, p2 *collection.Object) bool {
	s1 := (*p1).(Teacher)
	s2 := (*p2).(Teacher)
	return s1.Id < s2.Id
}

var t2 collection.Object
var t4 collection.Object
var t6 collection.Object
var t8 collection.Object
var st11 collection.Object
var st12 collection.Object
var st2 collection.Object
var st3 collection.Object
var st4 collection.Object
var s1 collection.Object
var s2 collection.Object
var s3 collection.Object
var s4 collection.Object
var kt reflect.Type
var vt reflect.Type
var m *SkipListMap

func RestartSkipListMap() {
	t2 = Teacher{2, "t2", 0}
	t4 = Teacher{4, "t4", 0}
	t6 = Teacher{6, "t6", 0}
	t8 = Teacher{8, "t8", 0}
	st11 = Student{100, "st11"}
	st12 = Student{101, "st12"}
	st2 = Student{102, "st2"}
	st3 = Student{103, "st3"}
	st4 = Student{104, "st4"}
	s1 = *set.New(reflect.TypeOf(Student{}))
	s2 = *set.New(reflect.TypeOf(Student{}))
	s3 = *set.New(reflect.TypeOf(Student{}))
	s4 = *set.New(reflect.TypeOf(Student{}))
	s1s := s1.(set.Set)
	s2s := s2.(set.Set)
	s3s := s3.(set.Set)
	s4s := s4.(set.Set)
	s1s.Add(&st11)
	s1s.Add(&st12)
	s2s.Add(&st2)
	s3s.Add(&st3)
	s4s.Add(&st4)
	kt = reflect.TypeOf(t2)
	vt = reflect.TypeOf(s1)
	m = New(kt, vt, precede)
	m.Put(&t2, &s1)
	m.Put(&t4, &s2)
	m.Put(&t6, &s3)
	m.Put(&t8, &s4)
}

func TestSetAll(t *testing.T) {
	TestSkipListMap_CeilingEntry(t)
	TestSkipListMap_Clear(t)
	TestSkipListMap_ContainsKey(t)
	TestSkipListMap_ContainsValue(t)
	TestSkipListMap_Empty(t)
	TestSkipListMap_EntrySet(t)
	TestSkipListMap_Equals(t)
	TestSkipListMap_FirstEntry(t)
	TestSkipListMap_FloorEntry(t)
	TestSkipListMap_Get(t)
	TestSkipListMap_GetEntryIterator(t)
	TestSkipListMap_HeadMap(t)
	TestSkipListMap_HigherEntry(t)
	TestSkipListMap_KeySet(t)
	TestSkipListMap_LastEntry(t)
	TestSkipListMap_LowerEntry(t)
	TestSkipListMap_PollFirstEntry(t)
	TestSkipListMap_PollLastEntry(t)
	TestSkipListMap_Put(t)
	TestSkipListMap_PutAll(t)
	TestSkipListMap_Remove(t)
	TestSkipListMap_SortedKeySet(t)
	TestSkipListMap_String(t)
	TestSkipListMap_SubMap(t)
	TestSkipListMap_TailMap(t)
	TestSkipListMap_Values(t)
}

func TestSkipListMap_CeilingEntry(t *testing.T) {
	RestartSkipListMap()
	var temp1 collection.Object = Teacher{4, "t4", 0}
	var temp2 collection.Object = Teacher{5, "t5", 0}
	var temp3 collection.Object = Teacher{8, "t8", 0}
	var temp4 collection.Object = Teacher{10, "t10", 0}
	k1 := (*m.CeilingEntry(&temp1)).GetKey()
	if *k1 != t4 {
		t.Error("CeilingEntry operation fail!")
	}
	k2 := (*m.CeilingEntry(&temp2)).GetKey()
	if *k2 != t6 {
		t.Error("CeilingEntry operation fail!")
	}
	k3 := (*m.CeilingEntry(&temp3)).GetKey()
	if *k3 != t8 {
		t.Error("CeilingEntry operation fail!")
	}
	k4 := m.CeilingEntry(&temp4)
	if k4 != nil {
		t.Error("CeilingEntry operation fail!")
	}
}

func TestSkipListMap_Clear(t *testing.T) {
	RestartSkipListMap()
	if m.Size() == 0 {
		t.Error("Start operation fail!")
	}
	m.Clear()
	if m.Size() != 0 {
		t.Error("Clear operation fail!")
	}
}

func TestSkipListMap_ContainsKey(t *testing.T) {
	RestartSkipListMap()
	var temp1 collection.Object = Teacher{4, "t4", 0}
	var temp2 collection.Object = Teacher{5, "t5", 0}
	if !m.ContainsKey(&temp1) {
		t.Error("ContainsKey operation fail!")
	}
	if m.ContainsKey(&temp2) {
		t.Error("ContainsKey operation fail!")
	}
	if m.ContainsKey(nil) {
		t.Error("ContainsKey operation fail!")
	}
}

func TestSkipListMap_ContainsValue(t *testing.T) {
	RestartSkipListMap()
	var st1t collection.Object = Student{100, "st11"}
	var st2t collection.Object = Student{101, "st12"}
	s1t := *set.New(reflect.TypeOf(Student{}))
	s1t.Add(&st1t)
	s1t.Add(&st2t)
	var s1to collection.Object = s1t
	if !m.ContainsValue(&s1to) {
		t.Error("ContainsValue operation fail!")
	}
}

func TestSkipListMap_Empty(t *testing.T) {
	RestartSkipListMap()
	if m.Empty() {
		t.Error("Start operation fail!")
	}
	m.Clear()
	if !m.Empty() {
		t.Error("Empty operation fail!")
	}
}

func TestSkipListMap_EntrySet(t *testing.T) {

}

func TestSkipListMap_Equals(t *testing.T) {
	RestartSkipListMap()
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
	ts1 = *set.New(reflect.TypeOf(Student{}))
	ts2 = *set.New(reflect.TypeOf(Student{}))
	ts3 = *set.New(reflect.TypeOf(Student{}))
	ts4 = *set.New(reflect.TypeOf(Student{}))
	ts1s := ts1.(set.Set)
	ts2s := ts2.(set.Set)
	ts3s := ts3.(set.Set)
	ts4s := ts4.(set.Set)
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
	var ma collection.Map = tm
	if !m.Equals(&ma) {
		t.Error("Equals operation fail!")
	}
}

func TestSkipListMap_FirstEntry(t *testing.T) {

}

func TestSkipListMap_FloorEntry(t *testing.T) {

}

func TestSkipListMap_Get(t *testing.T) {
	RestartSkipListMap()
	oj := m.Get(&t4)
	s := (*oj).(set.Set)
	if s.Size() != 1 || !s.Contains(&st2) {
		t.Error("Get operation fail!")
	}
}

func TestSkipListMap_GetEntryIterator(t *testing.T) {
	RestartSkipListMap()
	it := m.GetEntryIterator()
	index := 2
	s := ""
	for i := 0; it.HashNext(); i++ {
		if i == index {
			it.Remove()
		} else {
			entry := it.Next()
			teacher := (*(*entry).GetKey()).(Teacher)
			k := strconv.Itoa(teacher.Id)
			s += k
		}
	}
	if m.Size() != 3 || s != "2468" {
		t.Error("GetEntryIterator operation fail!")
	}
}

func TestSkipListMap_HeadMap(t *testing.T) {

}

func TestSkipListMap_HigherEntry(t *testing.T) {

}

func TestSkipListMap_KeySet(t *testing.T) {

}

func TestSkipListMap_LastEntry(t *testing.T) {

}

func TestSkipListMap_LowerEntry(t *testing.T) {

}

func TestSkipListMap_PollFirstEntry(t *testing.T) {

}

func TestSkipListMap_PollLastEntry(t *testing.T) {

}

func TestSkipListMap_Put(t *testing.T) {

}

func TestSkipListMap_PutAll(t *testing.T) {

}

func TestSkipListMap_Remove(t *testing.T) {

}

func TestSkipListMap_SortedKeySet(t *testing.T) {

}

func TestSkipListMap_String(t *testing.T) {

}

func TestSkipListMap_SubMap(t *testing.T) {

}

func TestSkipListMap_TailMap(t *testing.T) {

}

func TestSkipListMap_Values(t *testing.T) {

}