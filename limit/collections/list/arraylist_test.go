package list

import (
	"LimitGo/limit/collections"
	"reflect"
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

func TestNewArrayList(t *testing.T) {
	l := NewArrayList(reflect.TypeOf(Student{}))
	if l.GetType().Name() != "Student" || l.Size() != 0 ||
		len(l.elements) != 0 || cap(l.elements) != 10{
		t.Error("Create arraylist fail!")
	}
}

func TestArrayList_Append(t *testing.T) {
	l := NewArrayList(reflect.TypeOf(Student{}))
	var a collections.ListObject = Student{1, "Alice"}
	l.Append(&a)
	var b collections.ListObject = Teacher{1, "Bob", 0}
	l.Append(&b)
	if l.GetType().Name() != "Student" || l.Size() != 1 {
		t.Error("Add operation fail!")
	}
	if l.String() != "{{\"Id\":1,\"Name\":\"Alice\"}}" {
		t.Error("Add operation fail!")
	}
}

func TestArrayList_AddAll(t *testing.T) {
	l1 := NewArrayList(reflect.TypeOf(Student{}))
	var a collections.ListObject = Student{1, "Alice"}
	l1.Append(&a)
	l2 := NewArrayList(reflect.TypeOf(Student{}))
	var b collections.ListObject = Student{2, "Bob"}
	var c collections.ListObject = Student{3, "Mark"}
	l2.Append(&b)
	l2.Append(&c)
	if l1.Size() != 1 || l2.Size() != 2 {
		t.Error("Add operation fail!")
	}
	l1.AddAll(l2)
	if l1.Size() != 3 || l2.Size() != 2 {
		t.Error("AddAll operation fail!")
	}
	if l1.String() != "{{\"Id\":1,\"Name\":\"Alice\"},{\"Id\":2,\"Name\":\"Bob\"},{\"Id\":3,\"Name\":\"Mark\"}}" {
		t.Error("AddAll operation fail!")
	}
}

func TestArrayList_Clear(t *testing.T) {
	l := NewArrayList(reflect.TypeOf(Student{}))
	var a collections.ListObject = Student{1, "Alice"}
	var b collections.ListObject = Student{2, "Bob"}
	var c collections.ListObject = Student{3, "Mark"}
	l.Append(&a)
	l.Append(&b)
	l.Append(&c)
	l.Clear()
	if l.Size() != 0 {
		t.Error("Clear operation fail!")
	}
}

func TestArrayList_Contains(t *testing.T) {
	l := NewArrayList(reflect.TypeOf(Student{}))
	var a collections.ListObject = Student{1, "Alice"}
	var b collections.ListObject = Student{2, "Bob"}
	var c collections.ListObject = Student{3, "Mark"}
	var d collections.ListObject = Student{4, "Jessie"}
	l.Append(&a)
	l.Append(&b)
	l.Append(&c)
	if !l.Contains(&a) {
		t.Error("Contains operation fail!")
	}
	if l.Contains(&d) {
		t.Error("Contains operation fail!")
	}
}

func TestArrayList_Empty(t *testing.T) {
	l := NewArrayList(reflect.TypeOf(Student{}))
	if !l.Empty() {
		t.Error("Empty operation fail!")
	}
	var a collections.ListObject = Student{1, "Alice"}
	var b collections.ListObject = Student{2, "Bob"}
	l.Append(&a)
	l.Append(&b)
	if l.Empty() {
		t.Error("Empty operation fail!")
	}
}


func TestArrayList_Equals(t *testing.T) {
	l1 := NewArrayList(reflect.TypeOf(Student{}))
	l2 := NewArrayList(reflect.TypeOf(Student{}))
	l3 := NewArrayList(reflect.TypeOf(Student{}))
	var a collections.ListObject = Student{1, "Alice"}
	var b collections.ListObject = Student{2, "Bob"}
	var c collections.ListObject = Student{3, "Mark"}
	var d collections.ListObject = Student{4, "Jessie"}
	l1.Append(&a)
	l1.Append(&b)
	l2.Append(&a)
	l2.Append(&b)
	l3.Append(&c)
	l3.Append(&d)
	if !l1.Equals(l2) {
		t.Error("Equals operation fail!")
	}
	if l1.Equals(l3) {
		t.Error("Equals operation fail!")
	}
}