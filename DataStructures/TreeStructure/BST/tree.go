package bst

import (
	"errors"
	"fmt"
)

type BSTNode struct {
	Value int
	Left  *BSTNode
	Right *BSTNode
}

func (node *BSTNode) Append(value int) error {
	if value == node.Value {
		return errors.New("already exists")
	}
	if value > node.Value {
		if node.Right == nil {
			node.Right = &BSTNode{Value: value, Left: nil, Right: nil}
			return nil
		} else {
			node.Right.Append(value)
		}
	}
	if value < node.Value {
		if node.Left == nil {
			node.Left = &BSTNode{Value: value, Left: nil, Right: nil}
			return nil
		} else {
			node.Left.Append(value)
		}
	}
	return nil
}

func (node *BSTNode) FindMin() int {
	if node.Left == nil {
		return node.Value
	}
	return node.Left.FindMin()
}

func (node *BSTNode) FindMax() int {
	if node.Right == nil {
		return node.Value
	}
	return node.Right.FindMax()
}

func (node *BSTNode) PrintTree() {
	if node == nil {
		return
	}
	node.Left.PrintTree()
	fmt.Printf("%d ", node.Value)
	node.Right.PrintTree()
}
