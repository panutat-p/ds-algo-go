package singly_linked_list

type Node struct {
	Key      int
	NextNode *Node
}

func NewNode(data int) Node {
	return Node{
		Key:      data,
		NextNode: nil,
	}
}

// SetNextNode
// no need to implement setter because NextNode is public field
func (n *Node) SetNextNode(next *Node) {
	n.NextNode = next
}
