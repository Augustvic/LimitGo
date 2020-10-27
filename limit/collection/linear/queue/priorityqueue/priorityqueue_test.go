package priorityqueue

import (
	"LimitGo/limit/collection"
	"reflect"
	"testing"
)

type Student struct {
	Id int
	Name string
}

func TestArrayListAll(t *testing.T) {
	TestNew(t)
}

func TestNew(t *testing.T) {
	precede := func(p1 *collection.Object, p2 *collection.Object) bool {
		s1 := (*p1).(Student)
		s2 := (*p2).(Student)
		return s1.Id < s2.Id
	}
	l := New(reflect.TypeOf(Student{}), precede)
	if l.GetType().Name() != "Student" || l.Size() != 0 {
		t.Error("Create PriorityQueue fail!")
	}
}
