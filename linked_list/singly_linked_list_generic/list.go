package singly_linked_list_generic

type LinkedList[T any] struct {
	size int
	head *Node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.head
}

func (l *LinkedList[T]) Add(value T) {
	node := &Node[T]{Value: value}
	if l.head == nil {
		l.head = node
	} else {
		current := l.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
	l.size += 1
}

func (l *LinkedList[T]) Remove() {
	if l.head == nil {
		return
	}
	l.head = l.head.Next
	l.size -= 1
}

func (l *LinkedList[T]) Size() int {
	return l.size
}
