package collection

type Object interface{}

type Entry interface {
	GetKey() *Object
	GetValue() *Object
	SetValue(p *Object)
}

type equals func(p1 *Object, p2 *Object) bool
type precede func(p1 *Object, p2 *Object) bool