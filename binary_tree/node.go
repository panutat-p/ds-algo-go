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

// Insert
// fill empty child from top to bottom
// after the root is full, fill the left child
// result in an unbalanced binary tree
func (n *Node) Insert(num int) {
	if n.Left == nil {
		n.Left = NewNode(num)
		return
	}

	if n.Right == nil {
		n.Right = NewNode(num)
		return
	}

	n.Left.Insert(num) // more elements always go to the left child
}

func (n *Node) InOrder() {
	if n == nil {
		return
	}

	n.Left.InOrder()
	fmt.Printf("%v -> ", n.Key)
	n.Right.InOrder()
}

func (n *Node) PreOrder() {
	if n == nil {
		return
	}

	fmt.Printf("%v -> ", n.Key)
	n.Left.PreOrder()
	n.Right.PreOrder()
}

// AppendInOrder
// in-order tree traversal
func (n *Node) AppendInOrder(sl *[]int) {
	if n == nil {
		return
	}

	n.Left.AppendInOrder(sl)
	*sl = append(*sl, n.Key)
	n.Right.AppendInOrder(sl)
}

func (n *Node) InOrderIterative() {
	var (
		stack = NewStack()
		curr  = n // root
	)
	for {
		if curr != nil {
			stack.Enqueue(curr)
			curr = curr.Left
		} else if len(stack) != 0 {
			curr = stack.Dequeue()
			fmt.Printf("%v -> ", curr.Key)
			curr = curr.Right
		} else {
			break
		}
	}
}

type Stack []*Node

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Enqueue(n *Node) {
	*s = append(*s, n)
}

func (s *Stack) Dequeue() *Node {
	last := (*s)[s.Size()-1]
	*s = (*s)[:s.Size()-1]
	return last
}
