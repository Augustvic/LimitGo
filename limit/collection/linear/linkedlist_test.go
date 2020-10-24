package linear

import (
	"LimitGo/limit/collection"
	"reflect"
	"testing"
)

func TestLinkedListAll(t *testing.T) {
	TestNewLinkedList(t)
	TestLinkedList_Append(t)
	TestLinkedList_AddAll(t)
	TestLinkedList_AddAllHead(t)
	TestLinkedList_AddFirst(t)
	TestLinkedList_AddLast(t)
	TestLinkedList_Clear(t)
	TestLinkedList_Contains(t)
	TestLinkedList_Empty(t)
	TestLinkedList_Equals(t)
	TestLinkedList_First(t)
	TestLinkedList_Get(t)
	TestLinkedList_GetIterator(t)
	TestLinkedList_GetLast(t)
	TestLinkedList_IndexOf(t)
	TestLinkedList_Insert(t)
	TestLinkedList_Peek(t)
	TestLinkedList_Poll(t)
	TestLinkedList_Pop(t)
	TestLinkedList_Push(t)
	TestLinkedList_Remove(t)
	TestLinkedList_RemoveFirst(t)
	TestLinkedList_RemoveAt(t)
	TestLinkedList_RemoveLast(t)
	TestLinkedList_Set(t)
}

func TestNewLinkedList(t *testing.T) {
	l := NewLinkedList(reflect.TypeOf(Student{}))
	if l.GetType().Name() != "Student" || l.Size() != 0 {
		t.Error("Create LinkedList fail!")
	}
}

func TestLinkedList_Append(t *testing.T) {
	l := NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	l.Append(&a)
	var b collection.LinearObject = Teacher{1, "Bob", 0}
	l.Append(&b)
	if l.GetType().Name() != "Student" || l.Size() != 1 {
		t.Error("Append operation fail!")
	}
	if l.String() != "{{\"Id\":1,\"Name\":\"Alice\"}}" {
		t.Error("Append operation fail!")
	}
}

func TestLinkedList_AddAll(t *testing.T) {
	l1 := NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	l1.Append(&a)
	l2 := NewArrayList(reflect.TypeOf(Student{}))
	var b collection.LinearObject = Student{2, "Bob"}
	var c collection.LinearObject = Student{3, "Mark"}
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

func TestLinkedList_AddAllHead(t *testing.T) {
	l1 := NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	l1.Append(&a)
	l2 := NewLinkedList(reflect.TypeOf(Student{}))
	var b collection.LinearObject = Student{2, "Bob"}
	var c collection.LinearObject = Student{3, "Mark"}
	l2.Append(&b)
	l2.Append(&c)
	if l1.Size() != 1 || l2.Size() != 2 {
		t.Error("Append operation fail!")
	}
	var ll2 collection.Linear = l2
	l1.AddAllHead(&ll2)
	if l1.Size() != 3 || l2.Size() != 2 {
		t.Error("AddAll operation fail!")
	}
	if l1.String() != "{{\"Id\":2,\"Name\":\"Bob\"},{\"Id\":3,\"Name\":\"Mark\"},{\"Id\":1,\"Name\":\"Alice\"}}" {
		t.Error("AddAll operation fail!")
	}
}

func TestLinkedList_AddFirst(t *testing.T) {
	l := NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	l.AddFirst(&a)
	if l.Size() != 1 || l.String() != "{{\"Id\":1,\"Name\":\"Alice\"}}" {
		t.Error("AddFirst operation fail!")
	}
	var b collection.LinearObject = Student{2, "Bob"}
	var c collection.LinearObject = Student{3, "Mark"}
	l.AddFirst(&b)
	l.AddFirst(&c)
	if l.Size() != 3 || l.String() != "{{\"Id\":3,\"Name\":\"Mark\"},{\"Id\":2,\"Name\":\"Bob\"},{\"Id\":1,\"Name\":\"Alice\"}}" {
		t.Error("AddFirst operation fail!")
	}
}

func TestLinkedList_AddLast(t *testing.T) {
	l := NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	l.AddLast(&a)
	if l.Size() != 1 || l.String() != "{{\"Id\":1,\"Name\":\"Alice\"}}" {
		t.Error("AddLast operation fail!")
	}
	var b collection.LinearObject = Student{2, "Bob"}
	var c collection.LinearObject = Student{3, "Mark"}
	l.AddLast(&b)
	l.AddLast(&c)
	if l.Size() != 3 || l.String() != "{{\"Id\":1,\"Name\":\"Alice\"},{\"Id\":2,\"Name\":\"Bob\"},{\"Id\":3,\"Name\":\"Mark\"}}" {
		t.Error("AddLast operation fail!")
	}
}

func TestLinkedList_Clear(t *testing.T) {
	l := NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	var b collection.LinearObject = Student{2, "Bob"}
	var c collection.LinearObject = Student{3, "Mark"}
	l.Append(&a)
	l.Append(&b)
	l.Append(&c)
	l.Clear()
	if l.Size() != 0 {
		t.Error("Clear operation fail!")
	}
}

func TestLinkedList_Contains(t *testing.T) {
	l := NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	var b collection.LinearObject = Student{2, "Bob"}
	var c collection.LinearObject = Student{3, "Mark"}
	var d collection.LinearObject = Student{4, "Jessie"}
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

func TestLinkedList_Empty(t *testing.T) {
	l := NewArrayList(reflect.TypeOf(Student{}))
	if !l.Empty() {
		t.Error("Empty operation fail!")
	}
	var a collection.LinearObject = Student{1, "Alice"}
	var b collection.LinearObject = Student{2, "Bob"}
	l.Append(&a)
	l.Append(&b)
	if l.Empty() {
		t.Error("Empty operation fail!")
	}
}

func TestLinkedList_Equals(t *testing.T) {
	l1 := NewLinkedList(reflect.TypeOf(Student{}))
	var l2 collection.List = NewLinkedList(reflect.TypeOf(Student{}))
	var l3 collection.List = NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	var b collection.LinearObject = Student{2, "Bob"}
	var c collection.LinearObject = Student{3, "Mark"}
	var d collection.LinearObject = Student{4, "Jessie"}
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

func TestLinkedList_First(t *testing.T) {
	l := NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	var b collection.LinearObject = Student{2, "Bob"}
	l.Append(&a)
	l.Append(&b)
	c := l.First()
	if *c != a || l.Size() != 2 {
		t.Error("First operation fail!")
	}
}

func TestLinkedList_Get(t *testing.T) {
	l := NewArrayList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	var b collection.LinearObject = Student{2, "Bob"}
	var c collection.LinearObject = Student{3, "Mark"}
	var d collection.LinearObject = Student{4, "Jessie"}
	l.Append(&a)
	l.Append(&b)
	l.Append(&c)
	l.Append(&d)
	p := l.Get(2)
	if *p != c || l.Size() != 4 {
		t.Error("Get operation fail!")
	}
}

func TestLinkedList_GetIterator(t *testing.T) {
	l := NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	var b collection.LinearObject = Student{2, "Bob"}
	var c collection.LinearObject = Student{3, "Mark"}
	var d collection.LinearObject = Student{4, "Jessie"}
	l.Append(&a)
	l.Append(&b)
	l.Append(&c)
	l.Append(&d)
	var s = ""
	it := l.GetIterator()
	for it.HashNext() {
		p := it.Next()
		stu := (*p).(Student)
		s += stu.Name
	}
	if s != "AliceBobMarkJessie" {
		t.Error("Iterator fail!")
	}
}

func TestLinkedList_GetLast(t *testing.T) {
	l := NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	var b collection.LinearObject = Student{2, "Bob"}
	l.Append(&a)
	l.Append(&b)
	c := l.GetLast()
	if *c != b || l.Size() != 2 {
		t.Error("GetLast operation fail!")
	}
}

func TestLinkedList_IndexOf(t *testing.T) {
	l := NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	var b collection.LinearObject = Student{2, "Bob"}
	var c collection.LinearObject = Student{3, "Mark"}
	var d collection.LinearObject = Student{4, "Jessie"}
	l.Append(&a)
	l.Append(&b)
	l.Append(&c)
	l.Append(&d)
	i := l.IndexOf(&c)
	if i != 2 {
		t.Error("IndexOf operation fail!")
	}
}

func TestLinkedList_Insert(t *testing.T) {
	l := NewArrayList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	var b collection.LinearObject = Student{2, "Bob"}
	var c collection.LinearObject = Student{3, "Mark"}
	var d collection.LinearObject = Student{4, "Jessie"}
	var e collection.LinearObject = Student{5, "Mary"}
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

func TestLinkedList_Peek(t *testing.T) {
	l := NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	var b collection.LinearObject = Student{2, "Bob"}
	l.Append(&a)
	l.Append(&b)
	c := l.Peek()
	if *c != b || l.Size() != 2 {
		t.Error("Peek operation fail!")
	}
}

// Queue
func TestLinkedList_Poll(t *testing.T) {
	l := NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	var b collection.LinearObject = Student{2, "Bob"}
	l.Append(&a)
	l.Append(&b)
	c := l.Poll()
	if *c != a || l.Size() != 1 {
		t.Error("Poll operation fail!")
	}
}

// Stack
func TestLinkedList_Pop(t *testing.T) {
	l := NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	var b collection.LinearObject = Student{2, "Bob"}
	l.Append(&a)
	l.Append(&b)
	c := l.Pop()
	if *c != b || l.Size() != 1 {
		t.Error("Pop operation fail!")
	}
}

func TestLinkedList_Push(t *testing.T) {
	l := NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	var b collection.LinearObject = Student{2, "Bob"}
	l.Push(&a)
	l.Push(&b)
	if l.String() != "{{\"Id\":1,\"Name\":\"Alice\"},{\"Id\":2,\"Name\":\"Bob\"}}" {
		t.Error("Push operation fail!")
	}
}

func TestLinkedList_Remove(t *testing.T) {
	l := NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	var b collection.LinearObject = Student{2, "Bob"}
	var c collection.LinearObject = Student{3, "Mark"}
	var d collection.LinearObject = Student{4, "Jessie"}
	l.Append(&a)
	l.Append(&b)
	l.Append(&c)
	l.Append(&d)
	l.Remove(&c)
	if l.IndexOf(&d) != 2 || l.Contains(&c) || l.Size() != 3 {
		t.Error("Remove operation fail!")
	}
}

func TestLinkedList_RemoveFirst(t *testing.T) {
	l := NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	var b collection.LinearObject = Student{2, "Bob"}
	var c collection.LinearObject = Student{3, "Mark"}
	var d collection.LinearObject = Student{4, "Jessie"}
	l.Append(&a)
	l.Append(&b)
	l.Append(&c)
	l.Append(&d)
	l.RemoveFirst()
	if l.IndexOf(&d) != 2 || l.Contains(&a) || l.Size() != 3 {
		t.Error("RemoveFirst operation fail!")
	}
}

func TestLinkedList_RemoveAt(t *testing.T) {
	l := NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	var b collection.LinearObject = Student{2, "Bob"}
	var c collection.LinearObject = Student{3, "Mark"}
	var d collection.LinearObject = Student{4, "Jessie"}
	l.Append(&a)
	l.Append(&b)
	l.Append(&c)
	l.Append(&d)
	l.RemoveAt(2)
	if l.IndexOf(&d) != 2 || l.Contains(&c) || l.Size() != 3 {
		t.Error("RemoveAt operation fail!")
	}
}

func TestLinkedList_RemoveLast(t *testing.T) {
	l := NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	var b collection.LinearObject = Student{2, "Bob"}
	var c collection.LinearObject = Student{3, "Mark"}
	var d collection.LinearObject = Student{4, "Jessie"}
	l.Append(&a)
	l.Append(&b)
	l.Append(&c)
	l.Append(&d)
	l.RemoveLast()
	if l.Contains(&d) || l.Size() != 3 {
		t.Error("RemoveLast operation fail!")
	}
}

func TestLinkedList_Set(t *testing.T) {
	l := NewLinkedList(reflect.TypeOf(Student{}))
	var a collection.LinearObject = Student{1, "Alice"}
	var b collection.LinearObject = Student{2, "Bob"}
	var c collection.LinearObject = Student{3, "Mark"}
	var d collection.LinearObject = Student{4, "Jessie"}
	l.Append(&a)
	l.Append(&b)
	l.Append(&c)
	l.Set(1, &d)
	if l.IndexOf(&d) != 1 || l.Contains(&b) || !l.Contains(&d) || l.Size() != 3 {
		t.Error("Set operation fail!")
	}
}