package linear

import (
	"LimitGo/limit/collection"
	"reflect"
	"testing"
)

func TestLinkedListAll(t *testing.T) {
	TestNewLinkedList(t)

}

func TestNewLinkedList(t *testing.T) {
	l := NewLinkedList(reflect.TypeOf(Student{}))
	if l.GetType().Name() != "Student" || l.Size() != 0 {
		t.Error("Create linkedlist fail!")
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

func TestLinkedList_Add(t *testing.T) {

}

func TestLinkedList_AddAll(t *testing.T) {

}

func TestLinkedList_AddAllHead(t *testing.T) {

}

func TestLinkedList_AddFirst(t *testing.T) {

}

func TestLinkedList_AddLast(t *testing.T) {

}

func TestLinkedList_Clear(t *testing.T) {

}

func TestLinkedList_Contains(t *testing.T) {

}

func TestLinkedList_Empty(t *testing.T) {

}

func TestLinkedList_Equals(t *testing.T) {

}

func TestLinkedList_First(t *testing.T) {

}

func TestLinkedList_Get(t *testing.T) {

}

func TestLinkedList_GetFirst(t *testing.T) {

}

func TestLinkedList_GetIterator(t *testing.T) {

}

func TestLinkedList_GetLast(t *testing.T) {

}

func TestLinkedList_GetType(t *testing.T) {

}

func TestLinkedList_IndexOf(t *testing.T) {

}

func TestLinkedList_Insert(t *testing.T) {

}

func TestLinkedList_Peek(t *testing.T) {

}

func TestLinkedList_Poll(t *testing.T) {

}

func TestLinkedList_Pop(t *testing.T) {

}

func TestLinkedList_Push(t *testing.T) {

}

func TestLinkedList_Remove(t *testing.T) {

}

func TestLinkedList_RemoveFirst(t *testing.T) {

}

func TestLinkedList_RemoveAt(t *testing.T) {

}

func TestLinkedList_RemoveLast(t *testing.T) {

}

func TestLinkedList_Set(t *testing.T) {

}

func TestLinkedList_Size(t *testing.T) {

}

func TestLinkedList_String(t *testing.T) {

}