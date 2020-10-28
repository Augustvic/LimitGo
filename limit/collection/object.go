package collection

type Object interface{}

type Entry interface {
	GetKey() *Object
	GetValue() *Object
	SetValue(p *Object)
	HashCode()
}