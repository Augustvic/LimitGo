package sortedset

import (
	"LimitGo/limit/collection"
	"LimitGo/limit/collection/linear/hashset"
	"LimitGo/limit/collection/map/skiplistmap"
	"strconv"
	"testing"
)

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
var s *SortedSet

func RestartSortedSet() {
	t2 = Teacher{2, "t2", 0}
	t4 = Teacher{4, "t4", 0}
	t6 = Teacher{6, "t6", 0}
	t8 = Teacher{8, "t8", 0}
	var m collection.SortedMap = skiplistmap.New(precede)
	s = New(&m)
	s.Add(&t2)
	s.Add(&t4)
	s.Add(&t6)
	s.Add(&t8)
}

func TestSortedSetAll(t *testing.T) {
	TestSortedSet_AddAll(t)
	TestSortedSet_Ceiling(t)
	TestSortedSet_Clear(t)
	TestSortedSet_Contains(t)
	TestSortedSet_Empty(t)
	TestSortedSet_Equals(t)
	TestSortedSet_First(t)
	TestSortedSet_Floor(t)
	TestSortedSet_GetIterator(t)
	TestSortedSet_Higher(t)
	TestSortedSet_Last(t)
	TestSortedSet_Lower(t)
	TestSortedSet_PollFirst(t)
	TestSortedSet_PollLast(t)
	TestSortedSet_Remove(t)
	TestSortedSet_RemoveAll(t)
	TestSortedSet_RetainAll(t)
}

func TestSortedSet_AddAll(t *testing.T) {
	RestartSortedSet()
	var t10 collection.Object = Teacher{10, "t10", 0}
	var t12 collection.Object = Teacher{12, "t12", 0}
	var hs collection.Linear = hashset.New()
	hs.Add(&t10)
	hs.Add(&t12)
	s.AddAll(&hs)
	if s.Size() != 6 || !s.Contains(&t2) || !s.Contains(&t10) {
		t.Error("AddAll operation fail!")
	}
}

func TestSortedSet_Ceiling(t *testing.T) {
	RestartSortedSet()
	var temp1 collection.Object = Teacher{1, "t1", 0}
	var temp2 collection.Object = Teacher{5, "t5", 0}
	var temp3 collection.Object = Teacher{8, "t8", 0}
	var temp4 collection.Object = Teacher{10, "t10", 0}
	tt1 := s.Ceiling(&temp1)
	if *tt1 != t2 {
		t.Error("Ceiling operation fail!")
	}
	tt2 := s.Ceiling(&temp2)
	if *tt2 != t6 {
		t.Error("Ceiling operation fail!")
	}
	tt3 := s.Ceiling(&temp3)
	if *tt3 != t8 {
		t.Error("Ceiling operation fail!")
	}
	tt4 := s.Ceiling(&temp4)
	if tt4 != nil {
		t.Error("Ceiling operation fail!")
	}
}

func TestSortedSet_Clear(t *testing.T) {
	RestartSortedSet()
	if s.Size() == 0 {
		t.Error("Start operation fail!")
	}
	s.Clear()
	if s.Size() != 0 || (*s.m).Size() != 0 {
		t.Error("Clear operation fail!")
	}
}

func TestSortedSet_Contains(t *testing.T) {
	RestartSortedSet()
	var temp1 collection.Object = Teacher{2, "t2", 0}
	var temp2 collection.Object = Teacher{4, "t4", 0}
	var temp3 collection.Object = Teacher{5, "t5", 0}
	if !s.Contains(&temp1) {
		t.Error("Contains operation fail!")
	}
	if !s.Contains(&temp2) {
		t.Error("Contains operation fail!")
	}
	if s.Contains(&temp3) {
		t.Error("Contains operation fail!")
	}
	if s.Contains(nil) {
		t.Error("Contains operation fail!")
	}
}

func TestSortedSet_Empty(t *testing.T) {
	RestartSortedSet()
	if s.Empty() {
		t.Error("Start operation fail!")
	}
	s.Clear()
	if !s.Empty() {
		t.Error("Empty operation fail!")
	}
}

func TestSortedSet_Equals(t *testing.T) {
	RestartSortedSet()
	var tt2 collection.Object = Teacher{2, "t2", 0}
	var tt4 collection.Object = Teacher{4, "t4", 0}
	var tt6 collection.Object = Teacher{6, "t6", 0}
	var tt8 collection.Object = Teacher{8, "t8", 0}
	var tks collection.Set = hashset.New()
	var tsks collection.Set = hashset.New()
	tks.Add(&tt2)
	tks.Add(&tt4)
	tks.Add(&tt6)
	tks.Add(&tt8)
	tsks.Add(&tt4)
	tsks.Add(&tt6)
	tsks.Add(&tt8)
	if !s.Equals(&tks) || s.Equals(&tsks) {
		t.Error("Equals operation fail!")
	}
}

func TestSortedSet_First(t *testing.T) {
	RestartSortedSet()
	k1 := s.First()
	if *k1 != t2 {
		t.Error("First operation fail!")
	}
}

func TestSortedSet_Floor(t *testing.T) {
	RestartSortedSet()
	var temp1 collection.Object = Teacher{1, "t1", 0}
	var temp2 collection.Object = Teacher{5, "t5", 0}
	var temp3 collection.Object = Teacher{8, "t8", 0}
	var temp4 collection.Object = Teacher{11, "t10", 0}
	sk1 := s.Floor(&temp1)
	if sk1 != nil {
		t.Error("Floor operation fail!")
	}
	sk2 := s.Floor(&temp2)
	if *sk2 != t4 {
		t.Error("Floor operation fail!")
	}
	sk3 := s.Floor(&temp3)
	if *sk3 != t8 {
		t.Error("Floor operation fail!")
	}
	sk4 := s.Floor(&temp4)
	if *sk4 != t8 {
		t.Error("Floor operation fail!")
	}
}

func TestSortedSet_GetIterator(t *testing.T) {
	RestartSortedSet()
	it := s.GetIterator()
	index := 2
	st := ""
	for i := 0; it.HashNext(); i++ {
		if i == index {
			it.Remove()
		} else {
			key := it.Next()
			teacher := (*key).(Teacher)
			k := strconv.Itoa(teacher.Id)
			st += k
		}
	}
	if s.Size() != 3 || (*s.m).Size() != 3 || st != "2468" {
		t.Error("GetIterator operation fail!")
	}
}

func TestSortedSet_Higher(t *testing.T) {
	RestartSortedSet()
	var temp1 collection.Object = Teacher{1, "t1", 0}
	var temp2 collection.Object = Teacher{5, "t5", 0}
	var temp3 collection.Object = Teacher{6, "t8", 0}
	var temp4 collection.Object = Teacher{10, "t10", 0}
	sk1 := s.Higher(&temp1)
	if *sk1 != t2 {
		t.Error("Higher operation fail!")
	}
	sk2 := s.Higher(&temp2)
	if *sk2 != t6 {
		t.Error("Higher operation fail!")
	}
	sk3 := s.Higher(&temp3)
	if *sk3 != t8 {
		t.Error("Higher operation fail!")
	}
	sk4 := s.Higher(&temp4)
	if sk4 != nil {
		t.Error("Higher operation fail!")
	}
}

func TestSortedSet_Last(t *testing.T) {
	RestartSortedSet()
	k1 := s.Last()
	if *k1 != t8 {
		t.Error("Last operation fail!")
	}
}

func TestSortedSet_Lower(t *testing.T) {
	RestartSortedSet()
	var temp1 collection.Object = Teacher{1, "t1", 0}
	var temp2 collection.Object = Teacher{5, "t5", 0}
	var temp3 collection.Object = Teacher{8, "t8", 0}
	var temp4 collection.Object = Teacher{11, "t10", 0}
	k1 := s.Lower(&temp1)
	if k1 != nil {
		t.Error("Lower operation fail!")
	}
	k2 := s.Lower(&temp2)
	if *k2 != t4 {
		t.Error("Lower operation fail!")
	}
	k3 := s.Lower(&temp3)
	if *k3 != t6 {
		t.Error("Lower operation fail!")
	}
	k4 := s.Lower(&temp4)
	if *k4 != t8 {
		t.Error("Lower operation fail!")
	}
}

func TestSortedSet_PollFirst(t *testing.T) {
	RestartSortedSet()
	k1 := s.PollFirst()
	if s.Size() != 3 || *k1 != t2 || *s.First() != t4 {
		t.Error("PollFirst operation fail!")
	}
}

func TestSortedSet_PollLast(t *testing.T) {
	RestartSortedSet()
	k1 := s.PollLast()
	if s.Size() != 3 || *k1 != t8 || *s.Last() != t6 {
		t.Error("PollLast operation fail!")
	}
}

func TestSortedSet_Remove(t *testing.T) {
	RestartSortedSet()
	s.Remove(&t2)
	s.Remove(&t4)
	k1 := s.First()
	if s.Size() != 2 || *k1 != t6 {
		t.Error("Remove operation fail!")
	}
}

func TestSortedSet_RemoveAll(t *testing.T) {
	RestartSortedSet()
	var s1 collection.Linear = hashset.New()
	s1.Add(&t2)
	s1.Add(&t4)
	s.RemoveAll(&s1)
	if s.Size() != 2 || s.Contains(&t2) || s.Contains(&t4) || !s.Contains(&t6){
		t.Error("RemoveAll operation fail!")
	}
}

func TestSortedSet_RetainAll(t *testing.T) {
	RestartSortedSet()
	var s1 collection.Linear = hashset.New()
	s1.Add(&t2)
	s1.Add(&t4)
	s.RetainAll(&s1)
	if s.Size() != 2 || !s.Contains(&t2) || !s.Contains(&t4) || s.Contains(&t6) {
		t.Error("RetainAll operation fail!")
	}
}