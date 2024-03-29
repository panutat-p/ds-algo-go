package doubly_linked_list

import (
	"errors"
	"fmt"
	"strings"
)

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (li *LinkedList) Print() {
	var forward strings.Builder
	current := li.Head
	for {
		if current == nil {
			forward.WriteString("nil")
			break
		}
		tmp := fmt.Sprintf("%v -> ", current.Key)
		forward.WriteString(tmp)
		current = current.Next
	}
	fmt.Println(forward.String())

	var backward strings.Builder
	current = li.Tail
	for {
		if current == nil {
			backward.WriteString("nil")
			break
		}
		tmp := fmt.Sprintf("%v -> ", current.Key)
		backward.WriteString(tmp)
		current = current.Prev
	}
	fmt.Println(backward.String())
}

func (li *LinkedList) IsEmpty() bool {
	return li.Size == 0
}

func (li *LinkedList) Append(num int) {
	newNode := &Node{num, nil, nil}

	if li.IsEmpty() {
		li.Head, li.Tail = newNode, newNode
		li.Size += 1
		return
	}

	li.Tail.Next = newNode
	newNode.Prev = li.Tail
	li.Tail = li.Tail.Next
	li.Size += 1
}

func (li *LinkedList) Prepend(num int) {
	newNode := &Node{num, nil, nil}

	if li.IsEmpty() {
		li.Head, li.Tail = newNode, newNode
		li.Size += 1
		return
	}

	li.Head.Prev = newNode
	newNode.Next = li.Head
	li.Head = li.Head.Prev
	li.Size += 1
}

func (li *LinkedList) Insert(index int, num int) bool {
	if index < 0 || index > li.Size {
		return false
	}

	if index == 0 {
		li.Prepend(num)
		return true
	}

	if index == li.Size {
		li.Append(num)
		return true
	}

	current := li.Head
	for i := 0; i < index-1; i += 1 { // traverse to Node before index
		current = current.Next
	}
	newNode := &Node{num, current, current.Next} // add newNode after the current Node
	current.Next.Prev = newNode
	current.Next = newNode
	li.Size += 1
	return true
}

func (li *LinkedList) Remove(num int) bool {
	if li.IsEmpty() {
		return false
	}

	current := li.Head
	for current != nil {
		if current.Key == num {
			if li.Size == 1 {
				li.Head, li.Tail = nil, nil
			} else if current == li.Head {
				li.Head = li.Head.Next
				li.Head.Prev.Next = nil
				li.Head.Prev = nil
			} else if current == li.Tail {
				li.Tail = li.Tail.Prev
				li.Tail.Next.Prev = nil
				li.Tail.Next = nil
			} else {
				current.Prev.Next = current.Next
				current.Next.Prev = current.Prev
			}
			li.Size -= 1
			return true
		}
		current = current.Next
	}
	return false
}

func (li *LinkedList) Pop(index int) (int, error) {
	if li.IsEmpty() {
		return -1, errors.New("linked list is already empty")
	}

	if index < 0 || index >= li.Size {
		return -1, errors.New("index out of range")
	}

	if index == 0 {
		data := li.Head.Key
		li.Head = li.Head.Next
		li.Head.Prev.Next = nil
		li.Head.Prev = nil
		li.Size -= 1
		return data, nil
	}

	if index == li.Size-1 {
		data := li.Tail.Key
		li.Tail = li.Tail.Prev
		li.Tail.Next.Prev = nil
		li.Tail.Next = nil
		li.Size -= 1
		return data, nil
	}

	if index < li.Size/2 {
		current := li.Head
		for i := 0; i < index-1; i += 1 {
			current = current.Next
		}
		data := current.Next.Key
		current.Next.Next.Prev = current
		current.Next = current.Next.Next
		return data, nil
	} else {
		current := li.Tail
		for i := li.Size - 1; i > index+1; i -= 1 {
			current = current.Prev
		}
		data := current.Prev.Key
		current.Prev.Prev.Next = current
		current.Prev = current.Prev.Prev
		return data, nil
	}
}
