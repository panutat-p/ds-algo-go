package singly_linked_list

import (
	"fmt"
)

// LinkedList
// https://www.golangprograms.com/golang-program-for-implementation-of-linked-list.html
type LinkedList struct {
	Head *Node
}

func (li *LinkedList) Traverse() {
	current := li.Head
	for {
		if current == nil {
			fmt.Print("nil\n")
			break
		}
		fmt.Printf("%+v -> ", current)
		current = current.NextNode
	}
}

func (li *LinkedList) IsEmpty() bool {
	return li.Head == nil
}

func (li *LinkedList) Size() int {
	current := li.Head
	count := 0
	for current != nil {
		count += 1
		current = current.NextNode // tail.NextNode is always nil
	}
	return count
}

// Prepend
// handle empty linked list
func (li *LinkedList) Prepend(newNode *Node) {
	if li.IsEmpty() {
		li.Head = newNode
		return
	}

	newNode.SetNextNode(li.Head)
	li.Head = newNode
}

// Append
// handle empty linked list
// handle linked list with size of 1
func (li *LinkedList) Append(newNode *Node) {
	if li.IsEmpty() {
		li.Head = newNode
		return
	}

	var previous *Node
	current := li.Head
	for {
		if current == nil {
			previous.NextNode = newNode
			return
		}
		previous = current
		current = current.NextNode
	}
}

func (li *LinkedList) Insert(newNode *Node, index int) {
	if index == 0 {
		li.Prepend(newNode)
		return
	}

	if index > li.Size() {
		panic("index out of bound")
	}

	if index == li.Size() {
		li.Append(newNode)
	}

	current := li.Head
	for i := 0; i < index-1; i += 1 {
		current = current.NextNode
	}
	left := current
	right := current.NextNode
	left.NextNode = newNode
	newNode.NextNode = right
}

func (li *LinkedList) Delete(num int) bool {
	if li.Size() == 0 {
		return false
	}

	var previous *Node
	fmt.Println("**********previous**********")
	fmt.Println(previous)
	current := li.Head
	for {
		if current == nil {
			return false
		}

		if current.Key == num {
			if li.Size() == 1 { // remove Head
				li.Head = nil
				return true
			} else if li.Head == current { // remove Head, assign new Head
				li.Head = current.NextNode
				return true
			} else { // remove node
				previous.NextNode = current.NextNode
				return true
			}
		}
		previous = current
		current = current.NextNode
	}
}

func (li *LinkedList) Search(num int) bool {
	current := li.Head
	for current != nil {
		if current.Key == num {
			return true
		} else {
			current = current.NextNode
		}
	}
	return false
}
