package tree

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) FindNode(value int) *TreeNode {
	var result *TreeNode = nil
	if node == nil {
		return nil
	}
	if node.Value == value {
		return node
	}
	if node.Left != nil {
		result = node.Left.FindNode(value)
	}
	if node.Right != nil && result == nil {
		result = node.Right.FindNode(value)
	}
	return result
}
