package main

import (
	"LimitGo/limit/collection"
	"unsafe"
)

type People interface {}

type Integer struct {
	equals func(unsafe.Pointer, unsafe.Pointer) bool
	value *collection.LinearObject
}

type Student struct {}

func main() {
}

