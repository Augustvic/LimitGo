package skiplistmap

import "LimitGo/limit/collection"

// Index nodes represent the levels of the skip list.
type Index struct {
	level int
	node  *Node
	down  *Index
	right *Index
}

// Node hold keys and values, and are singly linked in sorted order.
type Node struct {
	key   *collection.Object
	value *collection.Object
	next  *Node
}

func (node *Node) GetKey() *collection.Object {
	return node.key
}

func (node *Node) GetValue() *collection.Object {
	return node.value
}

func (node *Node) SetValue(p *collection.Object) {
	node.value = p
}