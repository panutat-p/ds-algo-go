package singly_linked_list

import (
	"fmt"
	"testing"
)

func TestNodeConstructor(t *testing.T) {
	n1 := NewNode(2)

	n2 := NewNode(10)

	fmt.Printf("n1: %+v\n", n1)
	fmt.Printf("n2: %+v\n", n2)
	fmt.Printf("&n2: %p\n", &n2)

	n1.SetNextNode(&n2)
	fmt.Printf("\nn1.NextNode: %p\n\n", n1.NextNode)

	fmt.Printf("n1: %+v\n", n1)
	fmt.Printf("n2: %+v\n", n2)

	if n1.NextNode != &n2 {
		t.Errorf("got %v, but expect %v", n1.NextNode, n2)
	}

	if n1.NextNode.Key != 10 {
		t.Errorf("got %v, but expect %v", n1.NextNode.Key, 10)
	}
}

func TestLinkedListWith2Nodes(t *testing.T) {
	n1 := NewNode(1)
	n2 := NewNode(2)

	li := LinkedList{&n1}
	n1.SetNextNode(&n2)
	fmt.Printf("li: %+v\n", li)

	if li.Size() != 2 {
		t.Errorf("got %v, but expect %v", li.Size(), 2)
	}

	if li.IsEmpty() != false {
		t.Errorf("got %v, but expect %v", li.IsEmpty(), true)
	}
}

func TestLinkedListPrepend(t *testing.T) {
	n1 := NewNode(2)
	n2 := NewNode(3)

	li := LinkedList{&n1}
	n1.SetNextNode(&n2)
	li.Traverse()

	newNode := NewNode(1)
	li.Prepend(&newNode)
	li.Traverse()

	if li.Size() != 3 {
		t.Errorf("got %v, but expect %v", li.Size(), 2)
	}
}

func TestLinkedListAppendToEmpty(t *testing.T) {
	li := LinkedList{nil}
	newNode := NewNode(100)
	li.Append(&newNode)
	li.Traverse()

	if li.Size() != 1 {
		t.Errorf("got %v, but expect %v", li.Size(), 1)
	}
}

func TestLinkedListAppend(t *testing.T) {
	n1 := NewNode(10)
	n2 := NewNode(20)
	n1.SetNextNode(&n2)
	n3 := NewNode(30)
	n2.SetNextNode(&n3)

	li := LinkedList{&n1}
	n4 := NewNode(40)
	li.Append(&n4)
	li.Traverse()

	if li.Size() != 4 {
		t.Errorf("got %v, but expect %v", li.Size(), 4)
	}
}

func TestLinkedListSearch(t *testing.T) {
	n1 := NewNode(10)
	n2 := NewNode(20)
	n1.SetNextNode(&n2)
	n3 := NewNode(30)
	n2.SetNextNode(&n3)
	n4 := NewNode(40)
	n3.SetNextNode(&n4)

	li := LinkedList{&n1}
	li.Traverse()

	if li.Size() != 4 {
		t.Errorf("got %v, but expect %v", li.Size(), 4)
	}

	if li.Search(40) != true {
		t.Errorf("got %v, but expect %v", li.Size(), true)
	}

	if li.Search(50) != false {
		t.Errorf("got %v, but expect %v", li.Search(50), false)
	}
}

func TestLinkedListInsert(t *testing.T) {
	n1 := NewNode(10)
	n2 := NewNode(20)
	n1.SetNextNode(&n2)
	n4 := NewNode(40)
	n2.SetNextNode(&n4)

	li := LinkedList{&n1}
	li.Traverse()

	n3 := NewNode(30)
	li.Insert(&n3, 2)
	li.Traverse()

	if li.Size() != 4 {
		t.Errorf("got %v, but expect %v", li.Size(), 4)
	}
}

func TestLinkedListDelete(t *testing.T) {
	n1 := NewNode(10)
	n2 := NewNode(20)
	n1.SetNextNode(&n2)
	n3 := NewNode(30)
	n2.SetNextNode(&n3)
	n4 := NewNode(40)
	n3.SetNextNode(&n4)

	li := LinkedList{&n1}
	li.Delete(20)
	li.Traverse()

	if li.Size() != 3 {
		t.Errorf("got %v, but expect %v", li.Size(), 3)
	}
}

func TestLinkedListDeleteAtHead(t *testing.T) {
	n1 := NewNode(10)
	n2 := NewNode(20)
	n1.SetNextNode(&n2)
	n3 := NewNode(30)
	n2.SetNextNode(&n3)
	n4 := NewNode(40)
	n3.SetNextNode(&n4)

	li := LinkedList{&n1}
	rst := li.Delete(10)
	li.Traverse()

	if rst != true {
		t.Errorf("got %v, but expect %v", rst, true)
	}

	if li.Size() != 3 {
		t.Errorf("got %v, but expect %v", li.Size(), 3)
	}
}

func TestLinkedListDeleteSingleNodeList(t *testing.T) {
	n1 := NewNode(10)

	li := LinkedList{&n1}
	rst := li.Delete(10)
	li.Traverse()

	if rst != true {
		t.Errorf("got %v, but expect %v", rst, true)
	}

	if li.Size() != 0 {
		t.Errorf("got %v, but expect %v", li.Size(), 0)
	}
}
