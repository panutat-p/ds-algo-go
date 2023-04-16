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

func TestTree_traversal(t *testing.T) {
	tree := NewTree(0)
	tree.Insert(1)
	tree.Insert(2)

	tree.Print()
}
