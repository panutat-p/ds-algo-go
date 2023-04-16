package binary_tree

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func NewNode(data int) *Node {
	return &Node{
		Key:   data,
		Left:  nil,
		Right: nil,
	}
}

func (n *Node) Insert(num int) {
	if n.Left == nil {
		n.Left = NewNode(num)
		return
	}

	if n.Right == nil {
		n.Right = NewNode(num)
		return
	}

	n.Left.Insert(num)
}

func (n *Node) Traverse() {
	fmt.Printf("%v -> ", n.Key)

	if n.Left != nil {
		n.Left.Traverse()
	}
	if n.Right != nil {
		n.Right.Traverse()
	}
}
