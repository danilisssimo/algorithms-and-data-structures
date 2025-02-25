package BFSTree

import (
	fifo "algorithms-and-data-structures/DataStructures/Stack/FIFO"
	tree "algorithms-and-data-structures/DataStructures/TreeStructure/Tree"
	"fmt"
)

type Node struct {
	*tree.Node
}

func (node *Node) FindNode(value int) *tree.Node {
	return findNode(node.Node, value)
}

func findNode(node *tree.Node, value int) *tree.Node {
	var result *tree.Node = nil
	queue := fifo.GetNewQueue(1)
	if node == nil {
		return result
	}
	queue.Enqueue(node)
	for queue.GetLength() > 0 {
		queueNode := queue.Dequeue().(*tree.Node)
		fmt.Println(queueNode.Value)
		if queueNode.Value == value {
			result = queueNode
			break
		}
		if queueNode.Left != nil {
			queue.Enqueue(queueNode.Left)
		}
		if queueNode.Right != nil {
			queue.Enqueue(queueNode.Right)
		}
	}
	return result
}
