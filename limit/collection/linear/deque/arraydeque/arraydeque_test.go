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
	if q.GetType().Name() != "Student" || q.Size() != 0 ||
		len(q.elements) != 8 || cap(q.elements) != 8 || q.elements[0] != nil {
		t.Error("Create ArrayDeque fail!")
	}
}

func TestArrayDeque_AddFirst(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	var a collection.Object = Student{1, "Alice"}
	q.AddFirst(&a)
	if q.Size() != 1 || q.String() != "{{\"Id\":1,\"Name\":\"Alice\"}}" {
		t.Error("AddFirst operation fail!")
	}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	q.AddFirst(&b)
	q.AddFirst(&c)
	if q.Size() != 3 || q.String() != "{{\"Id\":3,\"Name\":\"Mark\"},{\"Id\":2,\"Name\":\"Bob\"},{\"Id\":1,\"Name\":\"Alice\"}}" {
		t.Error("AddFirst operation fail!")
	}
}

func TestArrayDeque_AddLast(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	var a collection.Object = Student{1, "Alice"}
	q.AddLast(&a)
	if q.Size() != 1 || q.String() != "{{\"Id\":1,\"Name\":\"Alice\"}}" {
		t.Error("AddLast operation fail!")
	}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	q.AddLast(&b)
	q.AddLast(&c)
	if q.Size() != 3 || q.String() != "{{\"Id\":1,\"Name\":\"Alice\"},{\"Id\":2,\"Name\":\"Bob\"},{\"Id\":3,\"Name\":\"Mark\"}}" {
		t.Error("AddLast operation fail!")
	}
}

func TestArrayDeque_Clear(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	q.AddFirst(&a)
	q.AddLast(&b)
	q.AddFirst(&c)
	q.Clear()
	if q.Size() != 0 {
		t.Error("Clear operation fail!")
	}
}

func TestArrayDeque_Contains(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	q.AddFirst(&a)
	q.AddLast(&b)
	q.AddFirst(&c)
	if !q.Contains(&a) {
		t.Error("Contains operation fail!")
	}
	if q.Contains(&d) {
		t.Error("Contains operation fail!")
	}
}

func TestArrayDeque_Empty(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	if !q.Empty() {
		t.Error("Empty operation fail!")
	}
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	q.AddLast(&a)
	q.AddLast(&b)
	if q.Empty() {
		t.Error("Empty operation fail!")
	}
}

func TestArrayDeque_GetFirst(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	q.AddLast(&a)
	q.AddLast(&b)
	c := q.GetFirst()
	if *c != a || q.Size() != 2 {
		t.Error("First operation fail!")
	}
}

func TestArrayDeque_GetIterator(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	q.AddFirst(&a)
	q.AddFirst(&b)
	q.AddLast(&c)
	q.AddLast(&d)
	var s = ""
	it := q.GetIterator()
	for i := 0; it.HashNext(); i++ {
		p := it.Next()
		stu := (*p).(Student)
		s += stu.Name
		if i >= 2 {
			it.Remove()
		}
	}
	if q.Size() != 2 || s != "BobAliceMarkJessie" {
		t.Error("Iterator fail!")
	}
}

func TestArrayDeque_GetLast(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	q.AddLast(&a)
	q.AddLast(&b)
	c := q.GetLast()
	if *c != b || q.Size() != 2 {
		t.Error("GetLast operation fail!")
	}
}

func TestArrayDeque_RemoveFirst(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	q.AddLast(&a)
	q.AddLast(&b)
	q.AddLast(&c)
	q.AddLast(&d)
	q.RemoveFirst()
	if q.Contains(&a) || q.Size() != 3 {
		t.Error("RemoveFirst operation fail!")
	}
}

func TestArrayDeque_RemoveLast(t *testing.T) {
	q := New(reflect.TypeOf(Student{}))
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	q.AddLast(&a)
	q.AddLast(&b)
	q.AddLast(&c)
	q.AddLast(&d)
	q.RemoveLast()
	if q.Contains(&d) || q.Size() != 3 {
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
	q.AddLast(&d)
	q.AddLast(&e)
	q.AddLast(&f)
	q.AddLast(&g)
	q.AddLast(&h)
	q.AddLast(&i)
	q.AddFirst(&c)
	q.AddFirst(&b)
	q.AddFirst(&a)
	if q.Size() != 9 || len(q.elements) != 16 {
		t.Error("DoubleLen operation fail!")
	}
}