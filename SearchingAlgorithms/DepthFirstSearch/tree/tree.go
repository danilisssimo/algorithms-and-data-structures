package tree

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) FindNode(value int) bool {
	var result bool = false
	if node == nil {
		return false
	}
	if node.Value == value {
		return true
	}
	if node.Left != nil {
		result = node.Left.FindNode(value)
	}
	if node.Right != nil && !result {
		result = node.Right.FindNode(value)
	}
	return result
}
