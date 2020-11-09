package sortedset

import (
	"LimitGo/limit/collection"
	"LimitGo/limit/collection/map/skiplistmap"
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
	TestSortedSet_Add(t)
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

func TestSortedSet_Add(t *testing.T) {
	RestartSortedSet()
}

func TestSortedSet_AddAll(t *testing.T) {
	RestartSortedSet()
}

func TestSortedSet_Ceiling(t *testing.T) {
	RestartSortedSet()
}

func TestSortedSet_Clear(t *testing.T) {
	RestartSortedSet()
}

func TestSortedSet_Contains(t *testing.T) {
	RestartSortedSet()
}

func TestSortedSet_Empty(t *testing.T) {
	RestartSortedSet()
}

func TestSortedSet_Equals(t *testing.T) {
	RestartSortedSet()
}

func TestSortedSet_First(t *testing.T) {
	RestartSortedSet()
}

func TestSortedSet_Floor(t *testing.T) {
	RestartSortedSet()
}

func TestSortedSet_GetIterator(t *testing.T) {
	RestartSortedSet()
}

func TestSortedSet_Higher(t *testing.T) {
	RestartSortedSet()
}

func TestSortedSet_Last(t *testing.T) {
	RestartSortedSet()
}

func TestSortedSet_Lower(t *testing.T) {
	RestartSortedSet()
}

func TestSortedSet_PollFirst(t *testing.T) {
	RestartSortedSet()
}

func TestSortedSet_PollLast(t *testing.T) {
	RestartSortedSet()
}

func TestSortedSet_Remove(t *testing.T) {
	RestartSortedSet()
}

func TestSortedSet_RemoveAll(t *testing.T) {
	RestartSortedSet()
}

func TestSortedSet_RetainAll(t *testing.T) {
	RestartSortedSet()
}