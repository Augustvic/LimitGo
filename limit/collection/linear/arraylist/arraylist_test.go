package arraylist

import (
	"LimitGo/limit/collection"
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

func TestArrayListAll(t *testing.T) {
	TestNew(t)
	TestArrayList_Append(t)
	TestArrayList_AddAll(t)
	TestArrayList_Clear(t)
	TestArrayList_Contains(t)
	TestArrayList_Empty(t)
	TestArrayList_Equals(t)
	TestArrayList_Get(t)
	TestArrayList_IndexOf(t)
	TestArrayList_Insert(t)
	TestArrayList_Remove(t)
	TestArrayList_Set(t)
	TestArrayList_GetIterator(t)
	TestArrayList_IntType(t)
}

func TestNew(t *testing.T) {
	l := New()
	if l.Size() != 0 ||
		len(l.elements) != 0 || cap(l.elements) != 8 {
		t.Error("Create ArrayList fail!")
	}
}

func TestArrayList_Append(t *testing.T) {
	l := New()
	var a collection.Object = Student{1, "Alice"}
	l.Append(&a)
	if l.Size() != 1 {
		t.Error("Append operation fail!")
	}
	if l.String() != "{{\"Id\":1,\"Name\":\"Alice\"}}" {
		t.Error("Append operation fail!")
	}
}

func TestArrayList_AddAll(t *testing.T) {
	l1 := New()
	var a collection.Object = Student{1, "Alice"}
	l1.Append(&a)
	l2 := New()
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	l2.Append(&b)
	l2.Append(&c)
	if l1.Size() != 1 || l2.Size() != 2 {
		t.Error("Append operation fail!")
	}
	var ll2 collection.Linear = l2
	l1.AddAll(&ll2)
	if l1.Size() != 3 || l2.Size() != 2 {
		t.Error("AddAll operation fail!")
	}
	if l1.String() != "{{\"Id\":1,\"Name\":\"Alice\"},{\"Id\":2,\"Name\":\"Bob\"},{\"Id\":3,\"Name\":\"Mark\"}}" {
		t.Error("AddAll operation fail!")
	}
}

func TestArrayList_Clear(t *testing.T) {
	l := New()
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	l.Append(&a)
	l.Append(&b)
	l.Append(&c)
	l.Clear()
	if l.Size() != 0 {
		t.Error("Clear operation fail!")
	}
}

func TestArrayList_Contains(t *testing.T) {
	l := New()
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
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
	l := New()
	if !l.Empty() {
		t.Error("Empty operation fail!")
	}
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	l.Append(&a)
	l.Append(&b)
	if l.Empty() {
		t.Error("Empty operation fail!")
	}
}


func TestArrayList_Equals(t *testing.T) {
	l1 := New()
	var l2 collection.List = New()
	var l3 collection.List = New()
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	l1.Append(&a)
	l1.Append(&b)
	l2.Append(&a)
	l2.Append(&b)
	l3.Append(&c)
	l3.Append(&d)
	if !l1.Equals(&l2) {
		t.Error("Equals operation fail!")
	}
	if l1.Equals(&l3) {
		t.Error("Equals operation fail!")
	}
}

func TestArrayList_Get(t *testing.T) {
	l := New()
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	l.Append(&a)
	l.Append(&b)
	l.Append(&c)
	l.Append(&d)
	p := l.Get(2)
	s := (*p).(Student)
	if s != c {
		t.Error("Get operation fail!")
	}
}

func TestArrayList_IndexOf(t *testing.T) {
	l := New()
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	l.Append(&a)
	l.Append(&b)
	l.Append(&c)
	l.Append(&d)
	i := l.IndexOf(&c)
	if i != 2 {
		t.Error("IndexOf operation fail!")
	}
}

func TestArrayList_Insert(t *testing.T) {
	l := New()
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	var e collection.Object = Student{5, "Mary"}
	l.Append(&b)
	l.Append(&d)
	if l.IndexOf(&d) != 1 || l.Size() != 2 {
		t.Error("Append operation fail!")
	}
	l.Insert(1, &c)
	if l.IndexOf(&c) != 1 || l.IndexOf(&d) != 2 || l.Size() != 3 {
		t.Error("Insert operation fail!")
	}
	l.Insert(-1, &a)
	if l.IndexOf(&a) != 0 || l.IndexOf(&b) != 1 || l.Size() != 4 {
		t.Error("Insert operation fail!")
	}
	l.Insert(10, &e)
	if l.IndexOf(&e) != 4 || l.Size() != 5 {
		t.Error("Insert operation fail!")
	}
}

func TestArrayList_Remove(t *testing.T) {
	l := New()
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	l.Append(&a)
	l.Append(&b)
	l.Append(&c)
	l.Append(&d)
	l.Remove(&c)
	if l.IndexOf(&d) != 2 || l.Contains(&c) || l.Size() != 3 {
		t.Error("Remove operation fail!")
	}
}

func TestArrayList_Set(t *testing.T) {
	l := New()
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	l.Append(&a)
	l.Append(&b)
	l.Append(&c)
	l.Set(1, &d)
	if l.IndexOf(&d) != 1 || l.Contains(&b) || !l.Contains(&d) || l.Size() != 3 {
		t.Error("Set operation fail!")
	}
}

func TestArrayList_GetIterator(t *testing.T) {
	l := New()
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	l.Append(&a)
	l.Append(&b)
	l.Append(&c)
	l.Append(&d)
	var s = ""
	it := l.GetIterator()
	for i := 0; it.HashNext(); i++ {
		p := it.Next()
		stu := (*p).(Student)
		s += stu.Name
		if i >= 2 {
			it.Remove()
		}
	}
	if l.Size() != 2 || s != "AliceBobMarkJessie" {
		t.Error("Iterator fail!")
	}
}

func TestArrayList_IntType(t *testing.T) {
	l := New()
	var a collection.Object = 1
	var b collection.Object = 2
	var c collection.Object = 3
	l.Append(&a)
	l.Append(&b)
	l.Append(&c)
	p := l.Get(1)
	if (*p).(int) != 2 {
		t.Error("Int fail!")
	}
}