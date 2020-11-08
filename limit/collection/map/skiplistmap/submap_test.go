package skiplistmap

import (
	"LimitGo/limit/collection"
	"LimitGo/limit/collection/linear/hashset"
	"reflect"
	"strconv"
	"testing"
)

var sm *SubMap

func RestartSubMap() {
	RestartSkipListMap()
	s := m.SubMap(&t2, false, nil, false)
	sm = (*s).(*SubMap)
}

func TestSubMapAll(t *testing.T) {
	TestSubMap_CeilingEntry(t)
	TestSubMap_Clear(t)
	TestSubMap_ContainsKey(t)
	TestSubMap_ContainsValue(t)
	TestSubMap_Empty(t)
	TestSubMap_EntrySet(t)
	TestSubMap_Equals(t)
	TestSubMap_FirstEntry(t)
	TestSubMap_FloorEntry(t)
	TestSubMap_Get(t)
	TestSubMap_GetEntryIterator(t)
	TestSubMap_HeadMap(t)
	TestSubMap_HigherEntry(t)
	TestSubMap_KeySet(t)
	TestSubMap_LastEntry(t)
	TestSubMap_LowerEntry(t)
	TestSubMap_PollFirstEntry(t)
	TestSubMap_PollLastEntry(t)
	TestSubMap_Put(t)
	TestSubMap_PutAll(t)
	TestSubMap_Remove(t)
	TestSubMap_String(t)
	TestSubMap_SubMap(t)
	TestSubMap_TailMap(t)
	TestSubMap_Values(t)
}

func TestSubMap_CeilingEntry(t *testing.T) {
	RestartSubMap()
	var temp1 collection.Object = Teacher{1, "t1", 0}
	var temp2 collection.Object = Teacher{5, "t5", 0}
	var temp3 collection.Object = Teacher{8, "t8", 0}
	var temp4 collection.Object = Teacher{10, "t10", 0}
	k1 := (*sm.CeilingEntry(&temp1)).GetKey()
	if *k1 != t4 {
		t.Error("CeilingEntry operation fail!")
	}
	k2 := (*sm.CeilingEntry(&temp2)).GetKey()
	if *k2 != t6 {
		t.Error("CeilingEntry operation fail!")
	}
	k3 := (*sm.CeilingEntry(&temp3)).GetKey()
	if *k3 != t8 {
		t.Error("CeilingEntry operation fail!")
	}
	k4 := sm.CeilingEntry(&temp4)
	if k4 != nil {
		t.Error("CeilingEntry operation fail!")
	}
}

func TestSubMap_Clear(t *testing.T) {
	RestartSubMap()
	if sm.Size() == 0 {
		t.Error("Start operation fail!")
	}
	sm.Clear()
	if sm.Size() != 0 {
		t.Error("Clear operation fail!")
	}
	if m.Size() != 1 {
		t.Error("Clear operation fail!")
	}
}

func TestSubMap_ContainsKey(t *testing.T) {
	RestartSubMap()
	var temp1 collection.Object = Teacher{2, "t2", 0}
	var temp2 collection.Object = Teacher{4, "t4", 0}
	var temp3 collection.Object = Teacher{5, "t5", 0}
	if sm.ContainsKey(&temp1) {
		t.Error("ContainsKey operation fail!")
	}
	if !sm.ContainsKey(&temp2) {
		t.Error("ContainsKey operation fail!")
	}
	if sm.ContainsKey(&temp3) {
		t.Error("ContainsKey operation fail!")
	}
	if sm.ContainsKey(nil) {
		t.Error("ContainsKey operation fail!")
	}
}

func TestSubMap_ContainsValue(t *testing.T) {
	RestartSubMap()
	var st1t collection.Object = Student{100, "st11"}
	var st2t collection.Object = Student{101, "st12"}
	var st3t collection.Object = Student{104, "st4"}
	s1t := *hashset.New(reflect.TypeOf(Student{}))
	s2t := *hashset.New(reflect.TypeOf(Student{}))
	s1t.Add(&st1t)
	s1t.Add(&st2t)
	s2t.Add(&st3t)
	var s1to collection.Object = s1t
	var s2to collection.Object = s2t
	if sm.ContainsValue(&s1to) {
		t.Error("ContainsValue operation fail!")
	}
	if !sm.ContainsValue(&s2to) {
		t.Error("ContainsValue operation fail!")
	}
}

func TestSubMap_Empty(t *testing.T) {
	RestartSubMap()
	if sm.Empty() {
		t.Error("Start operation fail!")
	}
	sm.Clear()
	if !sm.Empty() {
		t.Error("Empty operation fail!")
	}
}

func TestSubMap_EntrySet(t *testing.T) {
	RestartSubMap()
	es := sm.EntrySet()
	if (*es).Size() != 3 {
		t.Error("EntrySet operation fail!")
	}
}

func TestSubMap_Equals(t *testing.T) {
	RestartSubMap()
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
	var tm2 *SkipListMap
	tt2 = Teacher{2, "t2", 0}
	tt4 = Teacher{4, "t4", 0}
	tt6 = Teacher{6, "t6", 0}
	tt8 = Teacher{8, "t8", 0}
	tst11 = Student{100, "st11"}
	tst12 = Student{101, "st12"}
	tst2 = Student{102, "st2"}
	tst3 = Student{103, "st3"}
	tst4 = Student{104, "st4"}
	ts1 = *hashset.New(reflect.TypeOf(Student{}))
	ts2 = *hashset.New(reflect.TypeOf(Student{}))
	ts3 = *hashset.New(reflect.TypeOf(Student{}))
	ts4 = *hashset.New(reflect.TypeOf(Student{}))
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
	tm2 = New(tkt, tvt, precede)
	tm.Put(&tt2, &ts1)
	tm.Put(&tt4, &ts2)
	tm.Put(&tt6, &ts3)
	tm.Put(&tt8, &ts4)
	tm2.Put(&tt4, &ts2)
	tm2.Put(&tt6, &ts3)
	tm2.Put(&tt8, &ts4)
	var ma collection.Map = tm
	var ma2 collection.Map = tm2
	if sm.Equals(&ma) {
		t.Error("Equals operation fail!")
	}
	if !sm.Equals(&ma2) {
		t.Error("Equals operation fail!")
	}
}

func TestSubMap_FirstEntry(t *testing.T) {
	RestartSubMap()
	key := (*sm.FirstEntry()).GetKey()
	if *key != t4 {
		t.Error("FirstEntry operation fail!")
	}
}

func TestSubMap_FloorEntry(t *testing.T) {
	RestartSubMap()
	var temp1 collection.Object = Teacher{1, "t1", 0}
	var temp2 collection.Object = Teacher{5, "t5", 0}
	var temp3 collection.Object = Teacher{8, "t8", 0}
	var temp4 collection.Object = Teacher{11, "t10", 0}
	k1 := sm.FloorEntry(&temp1)
	if k1 != nil {
		t.Error("FloorEntry operation fail!")
	}
	k2 := (*sm.FloorEntry(&temp2)).GetKey()
	if *k2 != t4 {
		t.Error("FloorEntry operation fail!")
	}
	k3 := (*sm.FloorEntry(&temp3)).GetKey()
	if *k3 != t8 {
		t.Error("FloorEntry operation fail!")
	}
	k4 := (*sm.FloorEntry(&temp4)).GetKey()
	if *k4 != t8 {
		t.Error("FloorEntry operation fail!")
	}
}

func TestSubMap_Get(t *testing.T) {
	RestartSubMap()
	oj := sm.Get(&t4)
	s := (*oj).(hashset.HashSet)
	if s.Size() != 1 || !s.Contains(&st2) {
		t.Error("Get operation fail!")
	}
	oj2 := sm.Get(&t2)
	if oj2 != nil {
		t.Error("Get operation fail!")
	}
}

func TestSubMap_GetEntryIterator(t *testing.T) {
	RestartSubMap()
	it := sm.GetEntryIterator()
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
	if sm.Size() != 2 || s != "468" {
		t.Error("GetEntryIterator operation fail!")
	}
}

func TestSubMap_HeadMap(t *testing.T) {
	RestartSubMap()
	tm := sm.HeadMap(&t6, true)
	if (*tm).Size() != 2 || (*(*(*tm).FirstEntry()).GetKey()) != t4 ||  (*(*(*tm).LastEntry()).GetKey()) != t6 {
		t.Error("HeadMap operation fail!")
	}
}

func TestSubMap_HigherEntry(t *testing.T) {
	RestartSubMap()
	var temp1 collection.Object = Teacher{2, "t2", 0}
	var temp2 collection.Object = Teacher{4, "t4", 0}
	var temp3 collection.Object = Teacher{8, "t8", 0}
	k1 := (*sm.HigherEntry(&temp1)).GetKey()
	if *k1 != t4 {
		t.Error("HigherEntry operation fail!")
	}
	k2 := (*sm.HigherEntry(&temp2)).GetKey()
	if *k2 != t6 {
		t.Error("HigherEntry operation fail!")
	}
	k3 := sm.HigherEntry(&temp3)
	if k3 != nil {
		t.Error("HigherEntry operation fail!")
	}
}

func TestSubMap_KeySet(t *testing.T) {
	RestartSubMap()
	ks := sm.KeySet()
	if (*ks).String() != "{{\"Id\":4,\"Name\":\"t4\",\"Sex\":0},{\"Id\":6,\"Name\":\"t6\",\"Sex\":0},{\"Id\":8,\"Name\":\"t8\",\"Sex\":0}}" {
		t.Error("KeySet operation fail!")
	}
}

func TestSubMap_LastEntry(t *testing.T) {
	RestartSubMap()
	key := (*m.LastEntry()).GetKey()
	if *key != t8 {
		t.Error("LastEntry operation fail!")
	}
}

func TestSubMap_LowerEntry(t *testing.T) {
	RestartSubMap()
	var temp1 collection.Object = Teacher{3, "t3", 0}
	var temp2 collection.Object = Teacher{5, "t5", 0}
	var temp3 collection.Object = Teacher{8, "t8", 0}
	var temp4 collection.Object = Teacher{11, "t10", 0}
	k1 := sm.LowerEntry(&temp1)
	if k1 != nil {
		t.Error("LowerEntry operation fail!")
	}
	k2 := (*sm.LowerEntry(&temp2)).GetKey()
	if *k2 != t4 {
		t.Error("LowerEntry operation fail!")
	}
	k3 := (*sm.LowerEntry(&temp3)).GetKey()
	if *k3 != t6 {
		t.Error("LowerEntry operation fail!")
	}
	k4 := (*sm.LowerEntry(&temp4)).GetKey()
	if *k4 != t8 {
		t.Error("LowerEntry operation fail!")
	}
}

func TestSubMap_PollFirstEntry(t *testing.T) {
	RestartSubMap()
	sm.PollFirstEntry()
	key := (*sm.FirstEntry()).GetKey()
	if sm.Size() != 2 || *key != t6 {
		t.Error("PollFirstEntry operation fail!")
	}
}

func TestSubMap_PollLastEntry(t *testing.T) {
	RestartSubMap()
	tm := sm.HeadMap(&t6, true)
	(*tm).PollLastEntry()
	key := (*(*tm).LastEntry()).GetKey()
	if (*tm).Size() != 1 || *key != t4 {
		t.Error("PollLastEntry operation fail!")
	}
}

func TestSubMap_Put(t *testing.T) {
	RestartSubMap()
	RestartSkipListMap()
	var t1 collection.Object = Teacher{1, "t1", 0}
	var t10 collection.Object = Teacher{10, "t10", 0}
	var st5 collection.Object = Student{102, "st2"}
	var st6 collection.Object = Student{103, "st3"}
	s5 := *hashset.New(reflect.TypeOf(Student{}))
	s6 := *hashset.New(reflect.TypeOf(Student{}))
	s5.Add(&st5)
	s6.Add(&st6)
	var s5o collection.Object = s5
	var s6o collection.Object = s6
	sm.Put(&t1, &s5o)
	sm.Put(&t10, &s6o)
	sm.Put(&t4, &s5o)
	it := sm.GetEntryIterator()
	s := ""
	for i := 0; it.HashNext(); i++ {
		entry := it.Next()
		teacher := (*(*entry).GetKey()).(Teacher)
		k := strconv.Itoa(teacher.Id)
		s += k
	}
	if sm.Size() != 4 || s != "46810" {
		t.Error("Put operation fail!")
	}
	v := *sm.Get(&t4)
	if !reflect.DeepEqual(v, s5) {
		t.Error("Put operation fail!")
	}
}

func TestSubMap_PutAll(t *testing.T) {
	RestartSubMap()
	var t5 collection.Object = Teacher{5, "t5", 0}
	var t10 collection.Object = Teacher{10, "t10", 0}
	var st5 collection.Object = Student{102, "st2"}
	var st6 collection.Object = Student{103, "st3"}
	s5 := *hashset.New(reflect.TypeOf(Student{}))
	s6 := *hashset.New(reflect.TypeOf(Student{}))
	s5.Add(&st5)
	s6.Add(&st6)
	var s5o collection.Object = s5
	var s6o collection.Object = s6
	var m2 collection.Map = New(kt, vt, precede)
	m2.Put(&t5, &s5o)
	m2.Put(&t10, &s6o)
	sm.PutAll(&m2)
	it := sm.GetEntryIterator()
	s := ""
	for i := 0; it.HashNext(); i++ {
		entry := it.Next()
		teacher := (*(*entry).GetKey()).(Teacher)
		k := strconv.Itoa(teacher.Id)
		s += k
	}
	if sm.Size() != 5 || s != "456810" {
		t.Error("PutAll operation fail!")
	}
}

func TestSubMap_Remove(t *testing.T) {
	RestartSubMap()
	sm.Remove(&t2)
	key := (*sm.FirstEntry()).GetKey()
	if sm.Size() != 3 || *key != t4 {
		t.Error("Remove operation fail!")
	}
	if m.Size() != 4 {
		t.Error("Remove operation fail!")
	}
	sm.Remove(&t4)
	key2 := (*sm.FirstEntry()).GetKey()
	if sm.Size() != 2 || *key2 != t6 {
		t.Error("Remove operation fail!")
	}
}

func TestSubMap_String(t *testing.T) {
	RestartSubMap()
	if sm.String() != "{{\"Id\":4,\"Name\":\"t4\",\"Sex\":0}={},{\"Id\":6,\"Name\":\"t6\",\"Sex\":0}={},{\"Id\":8,\"Name\":\"t8\",\"Sex\":0}={}}" {
		t.Error("String operation fail!")
	}
}

func TestSubMap_SubMap(t *testing.T) {
	RestartSubMap()
	tm := sm.SubMap(nil, false, &t6, false)
	if (*tm).Size() != 1 || (*(*(*tm).FirstEntry()).GetKey()) != t4 || (*(*(*tm).LastEntry()).GetKey()) != t4 {
		t.Error("SubMap operation fail!")
	}
}

func TestSubMap_TailMap(t *testing.T) {
	RestartSubMap()
	tm := sm.TailMap(&t4, true)
	if (*tm).Size() != 3 || (*(*(*tm).FirstEntry()).GetKey()) != t4 ||  (*(*(*tm).LastEntry()).GetKey()) != t8 {
		t.Error("TailMap operation fail!")
	}
}

func TestSubMap_Values(t *testing.T) {
	RestartSubMap()
	vs := sm.Values()
	if (*vs).Size() != 3 {
		t.Error("Values operation fail!")
	}
}