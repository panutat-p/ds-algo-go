package binary_tree

import (
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

func TestTree_print_in_order(t *testing.T) {
	tree := NewTree(6)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(1)
	tree.Insert(3)

	tree.Print() // 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> end
}
