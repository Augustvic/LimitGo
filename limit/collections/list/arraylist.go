package list

import (
	"LimitGo/limit"
	"LimitGo/limit/collections"
)

type Iterator collections.Iterator
type List collections.List

type ArrayList struct {
	elements []limit.Object
}

type ListItr struct {
	List ArrayList
	Cursor int
	LastRet int
}

func NewArrayList() *ArrayList {
	l := ArrayList{}
	l.elements = make([]limit.Object, 0, 10)
	return &l
}

func (l *ArrayList) Size() int{
	return len(l.elements)
}

func (l *ArrayList) Empty() bool {
	return len(l.elements) == 0
}

func (l *ArrayList) Contains(p *limit.Object) bool {
	for _, v := range l.elements {
		if v == *p {
			return true
		}
	}
	return false
}

func (l *ArrayList) Add(p *limit.Object) {
	l.elements = append(l.elements, *p)
}

func (l *ArrayList) Insert(index int, p *limit.Object) {
	copy(l.elements[index+1:], l.elements[index:])
	l.elements[index] = *p
}

func (l *ArrayList) AddAll(list List) {
	l.capacity(l.Size() + list.Size())
	it := list.GetIterator()
	for it.HashNext() {
		p := it.Next()
		l.Add(p)
	}
}

func (l *ArrayList) Remove(o *limit.Object) {

}


func (l *ArrayList) RemoveAt(index int) {

}

func (l *ArrayList) Clear() {

}

func (l *ArrayList) Equals(list List) bool {
	return false
}

func (l *ArrayList) String() string {
	return ""
}

func (l *ArrayList) Get(index int) *limit.Object {
	if index < 0 || index >= l.Size() {
		panic("Index out of range!")
	}
	return &l.elements[index]
}

func (l *ArrayList) Set(index int, o *limit.Object) {

}

func (l *ArrayList) IndexOf(o *limit.Object) int {
	return 0
}

func (l *ArrayList) GetIterator() Iterator{
	return &ListItr{*l, 0, -1}
}

func (it *ListItr) HashNext() bool {
	cursor := it
	size := it
	return cursor != size
}

func (it *ListItr) Next() *limit.Object {
	if it.HashNext() {
		lastRet := it.Cursor
		it.Cursor++
		return &it.List.elements[lastRet]
	}
	return nil
}

func (it *ListItr) Remove() {

}


func (l *ArrayList) capacity(size int) {
	if size <= cap(l.elements) {
		// There is room to grow.  Extend the slice.
		l.elements = l.elements[:size]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		newSize := size
		if newSize < 2*len(l.elements) {
			newSize = 2 * len(l.elements)
		}
		newElements := make([]limit.Object, size, newSize)
		copy(newElements, l.elements)
		l.elements = newElements
	}
}