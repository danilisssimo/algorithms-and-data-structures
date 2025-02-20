package tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (node *Node) GetHeight() int {
	var leftTreeHeight int = 1
	var rightTreeHeight int = 1

	if node.Left != nil {
		leftTreeHeight += node.Left.GetHeight()
	}
	if node.Right != nil {
		rightTreeHeight += node.Right.GetHeight()
	}

	return max(leftTreeHeight, rightTreeHeight)
}

func (node *Node) GetMatrix() {
	// var leftRes []int
	// var rightRes []int
	// if node.Left != nil {
	// 	leftRes = append(leftRes, node.Left.GetMatrix()[0]...)
	// } else {
	// 	return [][]int{{node.Value}}
	// }
	// if node.Right != nil {
	// 	rightRes = append(rightRes, node.Right.GetMatrix()[0]...)
	// } else {
	// 	return [][]int{{node.Value}}
	// }
}

func PrintTree(node *Node) {
	// var leftRes []Node
	// var rightRes []Node
	// if node.Left != nil {
	// 	leftRes = node.Left.PrintTree()
	// }
	// if node.Right != nil {
	// 	rightRes = node.Left.PrintTree()
	// }
	// var mergeRes []Node = make([]Node, len(leftRes)+len(rightRes))
	// mergeRes = append(mergeRes, leftRes[:]...)
	// mergeRes = append(mergeRes, rightRes[:]...)
	// fmt.Println(mergeRes)
	// return mergeRes
}
