package binary_tree

import "fmt"

type Tree struct {
	Root *Node
	Size int
}

func NewTree(data int) Tree {
	return Tree{
		Root: NewNode(data),
		Size: 1,
	}
}

func (t *Tree) Insert(num int) {
	if t.Size == 0 {
		t.Root = NewNode(num)
		return
	}

	t.Root.Insert(num)
	t.Size += 1
}

func (t *Tree) PrintInOrder() {
	t.Root.InOrder()
	fmt.Println()
}

func (t *Tree) PrintPreOrder() {
	t.Root.PreOrder()
	fmt.Println()
}

func (t *Tree) Slice() []int {
	var (
		sl = make([]int, 0, t.Size)
	)
	t.Root.AppendInOrder(&sl)
	return sl
}
