package binarytree

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) Append(value int) error {
	if value == node.Value {
		return errors.New("already exists")
	}
	if value > node.Value {
		if node.Right == nil {
			node.Right = &TreeNode{Value: value, Left: nil, Right: nil}
			return nil
		} else {
			node.Right.Append(value)
		}
	}
	if value < node.Value {
		if node.Left == nil {
			node.Left = &TreeNode{Value: value, Left: nil, Right: nil}
			return nil
		} else {
			node.Left.Append(value)
		}
	}
	return nil
}

func (node *TreeNode) FindMin() int {
	if node.Left == nil {
		return node.Value
	}
	return node.Left.FindMin()
}

func (node *TreeNode) FindMax() int {
	if node.Right == nil {
		return node.Value
	}
	return node.Right.FindMax()
}

func (node *TreeNode) PrintTree() {
	if node == nil {
		return
	}
	node.Left.PrintTree()
	fmt.Printf("%d ", node.Value)
	node.Right.PrintTree()
}
