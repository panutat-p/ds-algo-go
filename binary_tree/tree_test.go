package binary_tree

import (
	"reflect"
	"testing"
)

func TestTree_create(t *testing.T) {
	tree := NewTree(0)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)

	want := 6
	got := tree.Size
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestTree_PrintPreOrder(t *testing.T) {
	tree := NewTree(1)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(3)
	tree.Insert(6)
	tree.Insert(4)
	tree.Insert(5)

	tree.PrintPreOrder() // 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 ->
}

func TestTree_PrintInOrder(t *testing.T) {
	tree := NewTree(6)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(1)
	tree.Insert(3)

	tree.PrintInOrder() // 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 ->
}

func TestTree_Slice(t *testing.T) {
	tree := NewTree(6)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(1)
	tree.Insert(3)

	want := []int{1, 2, 3, 4, 5, 6, 7}
	got := tree.Slice()
	if !reflect.DeepEqual(want, got) {
		t.Error("want", want, "but got", got)
	}
}
