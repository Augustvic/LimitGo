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
	if l.GetType().Name() != "Student" || Size() != 0 {
		t.Error("Create PriorityQueue fail!")
	}
}

func TestPriorityQueue_Clear(t *testing.T) {
	q := New(reflect.TypeOf(Student{}), precede)
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	Add(&a)
	Add(&b)
	Add(&c)
	Clear()
	if Size() != 0 {
		t.Error("Clear operation fail!")
	}
}

func TestPriorityQueue_Contains(t *testing.T) {
	q := New(reflect.TypeOf(Student{}), precede)
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	Add(&a)
	Add(&b)
	Add(&c)
	if !Contains(&a) {
		t.Error("Contains operation fail!")
	}
	if Contains(&d) {
		t.Error("Contains operation fail!")
	}
}

func TestPriorityQueue_Empty(t *testing.T) {
	q := New(reflect.TypeOf(Student{}), precede)
	if !Empty() {
		t.Error("Empty operation fail!")
	}
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	Add(&a)
	Add(&b)
	if Empty() {
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
	Add(&i)
	Add(&h)
	if First() != &h {
		t.Error("First operation fail!")
	}
	Add(&e)
	Add(&f)
	Add(&g)
	if First() != &e {
		t.Error("First operation fail!")
	}
	Add(&b)
	Add(&c)
	Add(&a)
	Add(&d)
	if First() != &a {
		t.Error("First operation fail!")
	}
}

func TestPriorityQueue_GetIterator(t *testing.T) {
	q := New(reflect.TypeOf(Student{}), precede)
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	Add(&a)
	Add(&b)
	Add(&c)
	Add(&d)
	it := GetIterator()
	for i := 0; it.HashNext(); i++ {
		if i >= 2 {
			it.Remove()
		}
		it.Next()
	}
	if Size() != 2 {
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
	Add(&f)
	Add(&d)
	Add(&e)
	Add(&b)
	Add(&c)
	Add(&a)
	if Poll() != &a {
		t.Error("Poll operation fail!")
	}
	if Size() != 5 {
		t.Error("Poll operation fail!")
	}
}

func TestPriorityQueue_String(t *testing.T) {
	q := New(reflect.TypeOf(Student{}), precede)
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	Add(&b)
	Add(&a)
	if String() != "{{\"Id\":1,\"Name\":\"Alice\"},{\"Id\":2,\"Name\":\"Bob\"}}" {
		t.Error("String operation fail!")
	}
}