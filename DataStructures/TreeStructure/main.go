package main

import (
	"fmt"
	"main/bst"
)

func main() {
	var HeadNode bst.BSTNode = bst.BSTNode{Value: 0, Right: nil, Left: nil}
	for index := -10; index < 11; index += 2 {
		err := HeadNode.Append(index)
		if err != nil {
			fmt.Printf("Error: '%s' to add %d node\n", err, index)
		}
		fmt.Printf("\nNow tree is - [ ")
		HeadNode.PrintTree()
		fmt.Println("]")
		fmt.Printf("Max: %d, Min: %d\n", HeadNode.FindMax(), HeadNode.FindMin())
	}
}
