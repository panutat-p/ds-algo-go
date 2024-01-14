package singly_linked_list_generic

type Node[T any] struct {
	Value T
	Next  *Node[T]
}
