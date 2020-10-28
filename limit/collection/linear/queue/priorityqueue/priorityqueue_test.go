package priorityqueue

import (
	"LimitGo/limit/collection"
	"reflect"
	"testing"
)

var precede = func(p1 *collection.Object, p2 *collection.Object) bool {
	s1 := (*p1).(Student)
	s2 := (*p2).(Student)
	return s1.Id < s2.Id
}

type Student struct {
	Id int
	Name string
}

func TestArrayListAll(t *testing.T) {
	TestNew(t)
	TestPriorityQueue_Clear(t)
	TestPriorityQueue_Contains(t)
	TestPriorityQueue_Empty(t)
	TestPriorityQueue_First(t)
	TestPriorityQueue_GetIterator(t)
	TestPriorityQueue_Poll(t)
	TestPriorityQueue_String(t)
}

func TestNew(t *testing.T) {
	l := New(reflect.TypeOf(Student{}), precede)
	if l.GetType().Name() != "Student" || l.Size() != 0 {
		t.Error("Create PriorityQueue fail!")
	}
}

func TestPriorityQueue_Clear(t *testing.T) {
	q := New(reflect.TypeOf(Student{}), precede)
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	q.Add(&a)
	q.Add(&b)
	q.Add(&c)
	q.Clear()
	if q.Size() != 0 {
		t.Error("Clear operation fail!")
	}
}

func TestPriorityQueue_Contains(t *testing.T) {
	q := New(reflect.TypeOf(Student{}), precede)
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	q.Add(&a)
	q.Add(&b)
	q.Add(&c)
	if !q.Contains(&a) {
		t.Error("Contains operation fail!")
	}
	if q.Contains(&d) {
		t.Error("Contains operation fail!")
	}
}

func TestPriorityQueue_Empty(t *testing.T) {
	q := New(reflect.TypeOf(Student{}), precede)
	if !q.Empty() {
		t.Error("Empty operation fail!")
	}
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	q.Add(&a)
	q.Add(&b)
	if q.Empty() {
		t.Error("Empty operation fail!")
	}
}

func TestPriorityQueue_First(t *testing.T) {
	q := New(reflect.TypeOf(Student{}), precede)
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	var e collection.Object = Student{5, "Alex"}
	var f collection.Object = Student{6, "Ellen"}
	var g collection.Object = Student{7, "August"}
	var h collection.Object = Student{8, "Jeff"}
	var i collection.Object = Student{9, "Jerry"}
	q.Add(&i)
	q.Add(&h)
	if q.First() != &h {
		t.Error("First operation fail!")
	}
	q.Add(&e)
	q.Add(&f)
	q.Add(&g)
	if q.First() != &e {
		t.Error("First operation fail!")
	}
	q.Add(&b)
	q.Add(&c)
	q.Add(&a)
	q.Add(&d)
	if q.First() != &a {
		t.Error("First operation fail!")
	}
}

func TestPriorityQueue_GetIterator(t *testing.T) {
	q := New(reflect.TypeOf(Student{}), precede)
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	q.Add(&a)
	q.Add(&b)
	q.Add(&c)
	q.Add(&d)
	it := q.GetIterator()
	for i := 0; it.HashNext(); i++ {
		if i >= 2 {
			it.Remove()
		}
		it.Next()
	}
	if q.Size() != 2 {
		t.Error("Iterator operation fail!")
	}
}

func TestPriorityQueue_Poll(t *testing.T) {
	q := New(reflect.TypeOf(Student{}), precede)
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	var e collection.Object = Student{5, "Alex"}
	var f collection.Object = Student{6, "Ellen"}
	q.Add(&f)
	q.Add(&d)
	q.Add(&e)
	q.Add(&b)
	q.Add(&c)
	q.Add(&a)
	if q.Poll() != &a {
		t.Error("Poll operation fail!")
	}
	if q.Size() != 5 {
		t.Error("Poll operation fail!")
	}
}

func TestPriorityQueue_String(t *testing.T) {
	q := New(reflect.TypeOf(Student{}), precede)
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	q.Add(&b)
	q.Add(&a)
	if q.String() != "{{\"Id\":1,\"Name\":\"Alice\"},{\"Id\":2,\"Name\":\"Bob\"}}" {
		t.Error("String operation fail!")
	}
}