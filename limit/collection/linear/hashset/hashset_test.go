package hashset

import (
	"LimitGo/limit/collection"
	"testing"
)

type Student struct {
	Id int
	Name string
}

func TestSetAll(t *testing.T) {
	TestNew(t)
	TestSet_Add(t)
	TestSet_AddAll(t)
	TestSet_Clear(t)
	TestSet_Contains(t)
	TestSet_Empty(t)
	TestSet_Equals(t)
	TestSet_GetIterator(t)
	TestSet_Remove(t)
	TestSet_RemoveAll(t)
	TestSet_RetainAll(t)
}

func TestNew(t *testing.T) {
	l := New()
	if l.Size() != 0 {
		t.Error("Create hashset fail!")
	}
}

func TestSet_Add(t *testing.T) {
	l := New()
	var a collection.Object = Student{1, "Alice"}
	l.Add(&a)
	if l.Size() != 1 {
		t.Error("Append operation fail!")
	}
	if l.String() != "{{\"Id\":1,\"Name\":\"Alice\"}}" {
		t.Error("Append operation fail!")
	}
}

func TestSet_AddAll(t *testing.T) {
	l1 := New()
	var a collection.Object = Student{1, "Alice"}
	l1.Add(&a)
	l2 := New()
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	l2.Add(&b)
	l2.Add(&c)
	if l1.Size() != 1 || l2.Size() != 2 {
		t.Error("Add operation fail!")
	}
	var ll2 collection.Linear = l2
	l1.AddAll(&ll2)
	if l1.Size() != 3 || l2.Size() != 2 {
		t.Error("AddAll operation fail!")
	}
	if !l1.Contains(&b) {
		t.Error("AddAll operation fail!")
	}
}

func TestSet_Clear(t *testing.T) {
	l := New()
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	l.Add(&a)
	l.Add(&b)
	l.Add(&c)
	l.Clear()
	if l.Size() != 0 {
		t.Error("Clear operation fail!")
	}
}

func TestSet_Contains(t *testing.T) {
	l := New()
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	l.Add(&a)
	l.Add(&b)
	l.Add(&c)
	if !l.Contains(&a) {
		t.Error("Contains operation fail!")
	}
	if l.Contains(&d) {
		t.Error("Contains operation fail!")
	}
}

func TestSet_Empty(t *testing.T) {
	l := New()
	if !l.Empty() {
		t.Error("Empty operation fail!")
	}
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	l.Add(&a)
	l.Add(&b)
	if l.Empty() {
		t.Error("Empty operation fail!")
	}
}

func TestSet_Equals(t *testing.T) {
	l1 := New()
	l2 := New()
	l3 := New()
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	l1.Add(&a)
	l1.Add(&b)
	l2.Add(&a)
	l2.Add(&b)
	l3.Add(&c)
	l3.Add(&d)
	var l2s collection.Set = l2
	var l3s collection.Set = l3
	if !l1.Equals(&l2s) {
		t.Error("Equals operation fail!")
	}
	if l1.Equals(&l3s) {
		t.Error("Equals operation fail!")
	}
}

func TestSet_GetIterator(t *testing.T) {
	l1 := New()
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	l1.Add(&a)
	l1.Add(&b)
	l1.Add(&c)
	l1.Add(&d)
	l2 := New()
	it := l1.GetIterator()
	for it.HashNext() {
		p := it.Next()
		l2.Add(p)
	}
	var l2s collection.Set = l2
	if !l1.Equals(&l2s) {
		t.Error("Iterator fail!")
	}
}

func TestSet_Remove(t *testing.T) {
	l := New()
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	l.Add(&a)
	l.Add(&b)
	l.Add(&c)
	l.Add(&d)
	l.Remove(&c)
	if l.Contains(&c) || l.Size() != 3 {
		t.Error("Remove operation fail!")
	}
}

func TestSet_RemoveAll(t *testing.T) {
	l1 := New()
	l2 := New()
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	l1.Add(&a)
	l1.Add(&b)
	l1.Add(&c)
	l1.Add(&d)
	l2.Add(&b)
	l2.Add(&c)
	l2.Add(&d)
	if l1.Size() != 4 || l2.Size() != 3 {
		t.Error("Add operation fail!")
	}
	var ll2 collection.Linear = l2
	l1.RemoveAll(&ll2)
	if l1.Size() != 1 || !l1.Contains(&a) || l1.Contains(&b) ||
		l1.Contains(&c) || l1.Contains(&d) {
		t.Error("RemoveAll operation fail!")
	}
}

func TestSet_RetainAll(t *testing.T) {
	l1 := New()
	l2 := New()
	var a collection.Object = Student{1, "Alice"}
	var b collection.Object = Student{2, "Bob"}
	var c collection.Object = Student{3, "Mark"}
	var d collection.Object = Student{4, "Jessie"}
	l1.Add(&a)
	l1.Add(&b)
	l1.Add(&c)
	l1.Add(&d)
	l2.Add(&b)
	l2.Add(&c)
	l2.Add(&d)
	if l1.Size() != 4 || l2.Size() != 3 {
		t.Error("Add operation fail!")
	}
	var ll2 collection.Linear = l2
	l1.RetainAll(&ll2)
	if l1.Size() != 3 || l1.Contains(&a) || !l1.Contains(&b) ||
		!l1.Contains(&c) || !l1.Contains(&d) {
		t.Error("RetainAll operation fail!")
	}
}