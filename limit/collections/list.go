package collections

import "LimitGo/limit"

type List interface {
	Size() int
	Empty() bool
	Contains(p *limit.Object) bool
	Add(p *limit.Object)
	Insert(index int, p *limit.Object)
	AddAll(list List)
	Remove(o *limit.Object)
	RemoveAt(index int)
	Clear()
	Equals(list List) bool
	String() string
	Get(index int) *limit.Object
	Set(index int, o *limit.Object)
	IndexOf(o *limit.Object) int
	GetIterator() Iterator
}