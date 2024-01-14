package doubly_linked_list

import (
	"testing"
)

func TestDoublyLinkedList_Insert(t *testing.T) {
	li := LinkedList{}
	li.Append(1)
	li.Append(2)
	li.Append(3)
	li.Prepend(0)
	li.Prepend(-1)
	li.Insert(1, 100)
	li.Print()

	if li.Size != 6 {
		t.Errorf("got %v, but expect %v", li.Size, 6)
	}
}

func TestDoublyLinkedList_Remove(t *testing.T) {
	li := LinkedList{}
	li.Append(1)
	li.Append(2)
	li.Append(3)
	li.Append(4)

	rst := li.Remove(3)
	li.Print()

	if !rst {
		t.Errorf("got %v, but expect %v", rst, true)
	}

	if li.Size != 3 {
		t.Errorf("got %v, but expect %v", li.Size, 3)
	}
}

func TestDoublyLinkedList_Remove_single_node(t *testing.T) {
	li := LinkedList{}
	li.Append(1)

	rst := li.Remove(1)
	li.Print()

	if !rst {
		t.Errorf("got %v, but expect %v", rst, true)
	}

	if li.Size != 0 {
		t.Errorf("got %v, but expect %v", li.Size, 0)
	}
}

func TestDoublyLinkedList_Pop(t *testing.T) {
	li := LinkedList{}
	li.Append(1)
	li.Append(2)
	li.Append(3)
	li.Append(4)

	rst, err := li.Pop(2)
	if err != nil {
		panic(err)
	}
	li.Print()

	if rst != 3 {
		t.Errorf("got %v, but expect %v", rst, 3)
	}
}
