package main

import (
	bst "algorithms-and-data-structures/DataStructures/TreeStructure/BST"
	tests "algorithms-and-data-structures/DataStructures/TreeStructure/Tree"
	tree "algorithms-and-data-structures/DataStructures/TreeStructure/Tree"
	"fmt"
)

func testForBST() {
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

func testForTree() {
	var arrOfTrees []*tree.Node = []*tree.Node{
		tests.Tree1,
		tests.Tree2,
		tests.Tree3,
		tests.Tree4,
		tests.Tree5,
		tests.Tree6,
		&tree.Node{Value: 1},
	}
	for _, headNode := range arrOfTrees {
		fmt.Printf("Tree: ")
		tree.PrintTree(headNode)
		fmt.Printf("\nHeight of node: %d\n", headNode.GetHeight())
	}
}

func main() {
	testForBST()
	testForTree()
}
