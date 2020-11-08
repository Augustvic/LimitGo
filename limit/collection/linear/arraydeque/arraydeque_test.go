package arraydeque

import (
	"LimitGo/limit/collection"
	"reflect"
	"testing"
)

type Student struct {
	Id int
	Name string
}

func TestArrayDequeAll(t *testing.T) {
	TestNew(t)
	TestArrayDeque_AddFirst(t)
	TestArrayDeque_AddLast(t)
	TestArrayDeque_Clear(t)
	TestArrayDeque_Contains(t)
	TestArrayDeque_Empty(t)
	TestArrayDeque_GetFirst(t)
	TestArrayDeque_GetIterator(t)
	TestArrayDeque_GetLast(t)
	TestArrayDeque_RemoveFirst(t)
	TestArrayDeque_RemoveLast(t)
	TestArrayDeque_DoubleLen(t)
}

func TestNew(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	if q.GetType().Name() != "Student" || Size() != 0 ||
		len(elements) != 8 || cap(elements) != 8 || elements[0] != nil {
		t.Error("Create ArrayDeque fail!")
	}
}

func TestArrayDeque_AddFirst(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	var a collection.Object = Student{1, "Alice"}
	AddFirst(&a)
	if Size() != 1 || String() != "{{\"Id\":1,\"Name\":\"Alice\"}}" {
		t.Error("AddFirst operation fail!")
	}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	AddFirst(&b)
	AddFirst(&c)
	if Size() != 3 || String() != "{{\"Id\":3,\"Name\":\"Mark\"},{\"Id\":2,\"Name\":\"Bob\"},{\"Id\":1,\"Name\":\"Alice\"}}" {
		t.Error("AddFirst operation fail!")
	}
}

func TestArrayDeque_AddLast(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	var a collection.Object = Student{1, "Alice"}
	AddLast(&a)
	if Size() != 1 || String() != "{{\"Id\":1,\"Name\":\"Alice\"}}" {
		t.Error("AddLast operation fail!")
	}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	AddLast(&b)
	AddLast(&c)
	if Size() != 3 || String() != "{{\"Id\":1,\"Name\":\"Alice\"},{\"Id\":2,\"Name\":\"Bob\"},{\"Id\":3,\"Name\":\"Mark\"}}" {
		t.Error("AddLast operation fail!")
	}
}

func TestArrayDeque_Clear(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	AddFirst(&a)
	AddLast(&b)
	AddFirst(&c)
	Clear()
	if Size() != 0 {
		t.Error("Clear operation fail!")
	}
}

func TestArrayDeque_Contains(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	AddFirst(&a)
	AddLast(&b)
	AddFirst(&c)
	if !Contains(&a) {
		t.Error("Contains operation fail!")
	}
	if Contains(&d) {
		t.Error("Contains operation fail!")
	}
}

func TestArrayDeque_Empty(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	if !Empty() {
		t.Error("Empty operation fail!")
	}
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	AddLast(&a)
	AddLast(&b)
	if Empty() {
		t.Error("Empty operation fail!")
	}
}

func TestArrayDeque_GetFirst(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	AddLast(&a)
	AddLast(&b)
	c := GetFirst()
	if *c != a || Size() != 2 {
		t.Error("First operation fail!")
	}
}

func TestArrayDeque_GetIterator(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	AddFirst(&a)
	AddFirst(&b)
	AddLast(&c)
	AddLast(&d)
	var s = ""
	it := GetIterator()
	for i := 0; it.HashNext(); i++ {
		p := it.Next()
		stu := (*p).(Student)
		s += stu.Name
		if i >= 2 {
			it.Remove()
		}
	}
	if Size() != 2 || s != "BobAliceMarkJessie" {
		t.Error("Iterator fail!")
	}
}

func TestArrayDeque_GetLast(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	AddLast(&a)
	AddLast(&b)
	c := GetLast()
	if *c != b || Size() != 2 {
		t.Error("GetLast operation fail!")
	}
}

func TestArrayDeque_RemoveFirst(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	AddLast(&a)
	AddLast(&b)
	AddLast(&c)
	AddLast(&d)
	RemoveFirst()
	if Contains(&a) || Size() != 3 {
		t.Error("RemoveFirst operation fail!")
	}
}

func TestArrayDeque_RemoveLast(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	AddLast(&a)
	AddLast(&b)
	AddLast(&c)
	AddLast(&d)
	RemoveLast()
	if Contains(&d) || Size() != 3 {
		t.Error("RemoveLast operation fail!")
	}
}

func TestArrayDeque_DoubleLen(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	var e collection.Object = Student{5, "Alex"}
	var f collection.Object = Student{6, "Ellen"}
	var g collection.Object = Student{7, "August"}
	var h collection.Object = Student{8, "Jeff"}
	var i collection.Object = Student{9, "Jerry"}
	AddLast(&d)
	AddLast(&e)
	AddLast(&f)
	AddLast(&g)
	AddLast(&h)
	AddLast(&i)
	AddFirst(&c)
	AddFirst(&b)
	AddFirst(&a)
	if Size() != 9 || len(elements) != 16 {
		t.Error("DoubleLen operation fail!")
	}
}