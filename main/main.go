package main

import (
	"LimitGo/limit/collection"
	"LimitGo/limit/collection/linear/arraylist"
	"fmt"
	"reflect"
)

type People struct {}

type Student struct {}

func main() {
	l := arraylist.New(reflect.TypeOf(People{}))
	var p collection.Object = People{}
	l.Append(&p)
	fmt.Print("sss")
}

