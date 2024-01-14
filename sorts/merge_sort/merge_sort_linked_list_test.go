package merge_sort

import (
	"fmt"
	singly_linked_list2 "github.com/panutat-p/fiset-complete-ds-go/linked_list/singly_linked_list"
	"testing"
)

func TestSplitLinkedListOdd(t *testing.T) {
	li := singly_linked_list2.LinkedList{}
	n3 := singly_linked_list2.Node{Key: 3}
	n2 := singly_linked_list2.Node{Key: 2, NextNode: &n3}
	n1 := singly_linked_list2.Node{Key: 1, NextNode: &n2}
	li.Head = &n1

	left, right := SplitLinkedList(li)
	left.Traverse()
	right.Traverse()

	if left.Size() != 1 {
		t.Errorf("got %v, but expect %v", left.Size(), 1)
	}

	if right.Size() != 2 {
		t.Errorf("got %v, but expect %v", right.Size(), 2)
	}
}

func TestSplitLinkedListEven(t *testing.T) {
	li := singly_linked_list2.LinkedList{}
	n4 := singly_linked_list2.Node{Key: 4}
	n3 := singly_linked_list2.Node{Key: 3, NextNode: &n4}
	n2 := singly_linked_list2.Node{Key: 2, NextNode: &n3}
	n1 := singly_linked_list2.Node{Key: 1, NextNode: &n2}
	li.Head = &n1
	li.Traverse()

	left, right := SplitLinkedList(li)
	fmt.Print("left: ")
	left.Traverse()
	fmt.Print("right: ")
	right.Traverse()

	if left.Size() != 2 {
		t.Errorf("got %v, but expect %v", left.Size(), 2)
	}

	if right.Size() != 2 {
		t.Errorf("got %v, but expect %v", right.Size(), 2)
	}
}

func TestMergeLinkedList(t *testing.T) {
	left := singly_linked_list2.LinkedList{}
	n3 := singly_linked_list2.Node{Key: 7}
	n2 := singly_linked_list2.Node{Key: 5, NextNode: &n3}
	n1 := singly_linked_list2.Node{Key: 1, NextNode: &n2}
	left.Head = &n1

	right := singly_linked_list2.LinkedList{}
	n33 := singly_linked_list2.Node{Key: 8}
	n22 := singly_linked_list2.Node{Key: 4, NextNode: &n33}
	n11 := singly_linked_list2.Node{Key: 2, NextNode: &n22}
	right.Head = &n11

	li := MergeLinkedList(left, right)
	li.Traverse()

	if li.Size() != 6 {
		t.Errorf("got %v, but expect %v", li.Size(), 6)
	}
}

func TestMergeSortLinkedList(t *testing.T) {
	li := singly_linked_list2.LinkedList{}
	n7 := singly_linked_list2.Node{Key: 1}
	n6 := singly_linked_list2.Node{Key: 50, NextNode: &n7}
	n5 := singly_linked_list2.Node{Key: 99, NextNode: &n6}
	n4 := singly_linked_list2.Node{Key: 9, NextNode: &n5}
	n3 := singly_linked_list2.Node{Key: 34, NextNode: &n4}
	n2 := singly_linked_list2.Node{Key: 2, NextNode: &n3}
	n1 := singly_linked_list2.Node{Key: 100, NextNode: &n2}
	li.Head = &n1
	li.Traverse()

	sortedList := MergeSortLinkedList(li)
	sortedList.Traverse()

	if sortedList.Size() != 7 {
		t.Errorf("got %v, but expect %v", sortedList.Size(), 7)
	}
}
