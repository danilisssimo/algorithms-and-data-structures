package tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (node *Node) FindNode(value int) {}
