package DFSTree

import (
	tree "algorithms-and-data-structures/DataStructures/TreeStructure/Tree"
)

type Node struct {
	*tree.Node
}

func (node *Node) FindNode(value int) *tree.Node {
	return findNode(node.Node, value)
}

func findNode(node *tree.Node, value int) *tree.Node {
	var result *tree.Node = nil
	if node == nil {
		return nil
	}
	if node.Value == value {
		return node
	}
	if node.Left != nil {
		result = findNode(node.Left, value)
	}
	if node.Right != nil && result == nil {
		result = findNode(node.Right, value)
	}
	return result
}
