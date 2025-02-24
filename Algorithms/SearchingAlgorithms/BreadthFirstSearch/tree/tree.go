package BFSTree

import (
	tree "algorithms-and-data-structures/DataStructures/TreeStructure/Tree"
)

type Node struct {
	*tree.Node
}

func (node *Node) FindNode(value int) *tree.Node {
	return findNode(node.Node, value)
}

func findNode(node *tree.Node, value int) *tree.Node { return node }
